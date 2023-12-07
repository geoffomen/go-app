package useraccountsrv

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strings"
	"time"

	"example.com/internal/app/common/base/entity"
	"example.com/internal/app/common/base/vo"
	"example.com/internal/app/example/useraccount/useraccountdm"
	"example.com/internal/pkg/myerr/myerrimp"
	"example.com/internal/pkg/uuidutil"
)

func (srv *Service) Register(ctx vo.SessionInfo, args useraccountdm.CreateRequestDto) (int64, error) {
	msgs, err := args.Validate()
	if err != nil {
		return 0, myerrimp.New(err).AddMsgf("%s", strings.Join(msgs, ";")).SetCode(http.StatusBadRequest)
	}

	salt := uuidutil.GenUuid()
	hashedPass := srv.GenPassword(args.Password, salt)

	accountEntity := useraccountdm.UseraccountEntity{
		BaseEntity: entity.BaseEntity{
			CreatedTime: time.Now(),
		},
		Account:  args.Account,
		Password: hashedPass,
		Salt:     salt,
		Name:     args.Account,
		Avatar:   "",
		Phone:    "",
	}
	accId, err := srv.repo.Create(*ctx.SetContext(context.Background()), accountEntity)

	if err != nil {
		return 0, err
	}
	return accId, nil
}

func (srv *Service) GenPassword(originPass, salt string) string {
	m5 := md5.New()
	m5.Write([]byte(originPass))
	m5.Write([]byte(salt))
	hashedPass := hex.EncodeToString(m5.Sum(nil))
	return hashedPass
}

func (srv *Service) Auth(req *http.Request) (*vo.SessionInfo, bool) {
	return &vo.SessionInfo{}, true
}
