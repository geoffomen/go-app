package userimp

import (
	"time"

	"github.com/geoffomen/go-app/examples/user"
	"github.com/geoffomen/go-app/internal/pkg/database"
	"github.com/geoffomen/go-app/internal/pkg/myerr"
	"github.com/geoffomen/go-app/internal/pkg/vo"
)

// Service ...
type Service struct {
	db *database.Client
}

var instance *Service = &Service{}

// New ...
func New(db *database.Client,
) *Service {
	*instance = Service{
		db: db,
	}
	db.GetStmt().AutoMigrate(&UserEntity{})
	return instance
}

// GetInstance ..
func GetInstance() *Service {
	return instance
}

// NewInstanceWithDBClient ..
func (srv *Service) NewWithDb(db *database.Client) user.Iface {
	newSrv := *instance
	newSrv.db = db
	return &newSrv
}

func (srv *Service) Create(param user.CreateUserRequestDto) (int, error) {
	userEntity := UserEntity{
		BaseEntity: vo.BaseEntity{
			CreatedTime: vo.Mytime{
				Time: time.Now(),
			},
			UpdatedTime: vo.Mytime{
				Time: time.Now(),
			},
			Version: 0,
		},
		Name:     param.Name,
		NickName: param.NickName,
		Avatar:   param.Avatar,
		Phone:    param.Phone,
	}
	err := srv.db.GetStmt().Table(userEntity.TableName()).Create(&userEntity)
	if err != nil {
		return 0, myerr.New(err)
	}

	return userEntity.Id, nil
}

func (srv *Service) GetUserInfo(param user.GetUserInfoRequestDto) (user.UserInfoResponseDto, error) {
	userEntity := UserEntity{}
	err := srv.db.GetStmt().
		Table(userEntity.TableName()).
		Where("id=?", param.Id).
		First(&userEntity)
	if err != nil {
		return user.UserInfoResponseDto{}, myerr.New(err)
	}

	rt := user.UserInfoResponseDto{
		Id:       userEntity.Id,
		Name:     userEntity.Name,
		NickName: userEntity.NickName,
	}

	return rt, nil
}

func (srv *Service) Page(param user.PageRequestDto) (*vo.PageResponseDto, error) {
	sql := srv.db.GetStmt().Table(UserEntity{}.TableName())
	if param.Keyword != "" {
		sql.Where("name like ?", "%"+param.Keyword+"%")
	}
	if param.Tm.Time.Year() > 1 {
		sql.Where("created_time > ?", param.Tm)
	}
	if param.Sort != "" {
		sql.Order(param.Sort)
	}

	var total int64
	if param.HasTotal == 1 {
		sql.Count(&total)
	}

	es := make([]UserEntity, 0)
	err := sql.Offset(param.Offset).Limit(param.PageSize).Find(&es)
	if err != nil {
		return nil, myerr.New(err)
	}
	dt := make([]user.PageResponseVo, 0, len(es))
	for _, item := range es {
		o := user.PageResponseVo{
			Id:   item.Id,
			Name: item.Name,
		}
		dt = append(dt, o)
	}
	return &vo.PageResponseDto{
		Total:    total,
		PageSize: len(es),
		List:     dt,
	}, nil
}
