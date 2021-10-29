package accountimp

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/geoffomen/go-app/examples/account"
	"github.com/geoffomen/go-app/examples/user"
	"github.com/geoffomen/go-app/internal/pkg/config"
	"github.com/geoffomen/go-app/internal/pkg/database"
	"github.com/geoffomen/go-app/internal/pkg/digestutil"
	"github.com/geoffomen/go-app/internal/pkg/myerr"
	"github.com/geoffomen/go-app/internal/pkg/vo"
)

// Service ...
type Service struct {
	db          *database.Client
	config      config.Iface
	userService user.Iface
}

var instance *Service = &Service{}

// New ...
func New(db *database.Client,
	config config.Iface,
	userService user.Iface,
) *Service {
	*instance = Service{
		db:          db,
		config:      config,
		userService: userService,
	}
	db.GetStmt().AutoMigrate(&AccountEntity{})
	db.GetStmt().AutoMigrate(&LoginTokenEntity{})
	return instance
}

// GetInstance ..
func GetInstance() *Service {
	return instance
}

// NewInstanceWithDBClient ..
func (srv Service) NewWithDb(cl *database.Client) account.Iface {
	newSrv := *instance
	newSrv.db = cl
	return &newSrv
}

func (srv *Service) Register(param account.CreateRequestDto) (int, error) {
	salt := digestutil.GenUuid()
	pass := digestutil.Md5Encryption(param.Password, salt)

	accountEntity := AccountEntity{
		BaseEntity: vo.BaseEntity{
			CreatedTime: vo.Mytime{
				Time: time.Now(),
			},
			UpdatedTime: vo.Mytime{
				Time: time.Now(),
			},
			Version: 0,
		},
		Account:  param.Account,
		Password: pass,
		Salt:     salt,
	}
	err := srv.db.GetStmt().DoTransaction(func(tx *database.Client) error {
		uid, err := srv.userService.NewWithDb(tx).Create(user.CreateUserRequestDto{
			Name:     param.Account,
			NickName: param.Account,
		})
		if err != nil {
			return err
		}
		accountEntity.Uid = uid
		err = tx.GetStmt().Table(accountEntity.TableName()).Create(&accountEntity)
		if err != nil {
			return myerr.New(err)
		}
		return nil
	})

	if err != nil {
		return 0, err
	}
	return accountEntity.Id, nil
}

//CreateToken Creating Access Token
func (srv Service) CreateToken(uid int) (string, error) {
	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["uid"] = uid
	atClaims["issueAt"] = time.Now()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenString, err := at.SignedString([]byte(srv.config.GetStringOrDefault("server.secret", "")))
	if err != nil {
		return "", myerr.New(err)
	}

	return tokenString, nil
}

// ValidAndGetTokenData ..
func (srv Service) ValidAndGetTokenData(tokenString string) (*vo.SessionInfo, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, myerr.NewfWithCode(http.StatusUnauthorized, "Unexpected signing method: %v", token.Method).
				AddMsgf("令牌无效，请重新登录")
		}
		return []byte(srv.config.GetStringOrDefault("server.secret", "")), nil
	})
	if err != nil {
		return nil, myerr.NewfWithCode(http.StatusUnauthorized, err.Error()).AddMsgf("令牌无效，请重新登录")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uid := int(claims["uid"].(float64))
		loginToken := &LoginTokenEntity{}
		err := srv.db.GetStmt().
			Table(loginToken.TableName()).
			Where("token = ?", tokenString).
			First(&loginToken)
		if err != nil {
			return nil, myerr.NewfWithCode(http.StatusUnauthorized, "%s", err).AddMsgf("令牌已失效，请重新登录")
		}
		if loginToken.ExpireAt < time.Now().Unix() {
			return nil, myerr.NewfWithCode(http.StatusUnauthorized, "令牌已失效，请重新登录")
		}
		return &vo.SessionInfo{
			Uid:           uid,
			Token:         tokenString,
			TokenExpireAt: loginToken.ExpireAt,
		}, nil
	}
	return nil, myerr.NewfWithCode(http.StatusUnauthorized, "令牌无效，请重新登录")
}

// Login ...
func (srv Service) Login(param account.LoginRequestDto) (*account.LoginResponseDto, error) {
	accEntity := AccountEntity{}
	err := srv.db.GetStmt().
		Table(accEntity.TableName()).
		Where("account = ?", param.Account).
		First(&accEntity)
	if err != nil {
		return nil, myerr.New(err)
	}
	if accEntity.Password != digestutil.Md5Encryption(param.Password, accEntity.Salt) {
		return nil, myerr.Newf("用户名或密码错误")
	}

	userInfo, err := srv.userService.GetUserInfo(user.GetUserInfoRequestDto{Id: accEntity.Uid})
	if err != nil {
		return nil, err
	}

	tokenString, err := srv.CreateToken(userInfo.Id)
	if err != nil {
		return nil, myerr.New(err)
	}
	err = srv.db.GetStmt().
		Where("uid = ?", accEntity.Uid).
		Delete(&LoginTokenEntity{})
	if err != nil {
		return nil, myerr.New(err)
	}
	loginToken := LoginTokenEntity{
		BaseEntity: vo.BaseEntity{
			CreatedTime: vo.Mytime{
				Time: time.Now(),
			},
			UpdatedTime: vo.Mytime{
				Time: time.Now(),
			},
			Version: 0,
		},
		UID:      accEntity.Uid,
		Token:    tokenString,
		ExpireAt: time.Now().Unix() + int64(60*60*24*30),
	}
	err = srv.db.GetStmt().
		Table(loginToken.TableName()).
		Create(&loginToken)
	if err != nil {
		return nil, myerr.New(err)
	}

	rt := account.LoginResponseDto{
		Uid:         userInfo.Id,
		IssueAt:     vo.Mytime{Time: time.Now()},
		ExpireAt:    vo.Mytime{Time: time.Unix(loginToken.ExpireAt, 0)},
		TokenType:   "Bearer",
		AccessToken: tokenString,
	}
	return &rt, nil
}

// Logout ..
func (srv *Service) Logout(sessData vo.SessionInfo) (int, error) {
	err := srv.db.GetStmt().
		Where("token = ?", sessData.Token).
		Delete(&LoginTokenEntity{})
	if err != nil {
		return 0, myerr.NewfWithCode(500, "%s", err)
	}
	return 0, nil
}
