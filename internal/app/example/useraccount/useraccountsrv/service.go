package useraccountsrv

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strings"
	"time"

	"ibingli.com/internal/app/example/useraccount/useraccountdm"
	"ibingli.com/internal/pkg/myDatabase"
	"ibingli.com/internal/pkg/myErr/myErrImp"
	"ibingli.com/internal/pkg/myHttpServer"
	"ibingli.com/internal/pkg/uuidUtil"
)

func (srv *Service) Register(ctx *myHttpServer.SessionInfo, args *useraccountdm.CreateRequestDto) (int64, error) {
	msgs, err := args.Validate()
	if err != nil {
		return 0, myErrImp.New(err).AddMsgf("%s", strings.Join(msgs, ";")).SetCode(http.StatusBadRequest)
	}

	salt := uuidUtil.GenUuid()
	hashedPass := srv.GenPassword(args.Password, salt)

	accountEntity := useraccountdm.UseraccountEntity{
		BaseEntity: myDatabase.BaseEntity{
			CreatedTime: time.Now(),
		},
		Account:  args.Account,
		Password: hashedPass,
		Salt:     salt,
		Name:     args.Account,
		Avatar:   "",
		Phone:    "",
	}
	accId, err := srv.repo.Create(ctx, srv.db, &accountEntity)

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

func (srv *Service) Auth(req *http.Request) (*myHttpServer.SessionInfo, bool) {
	return &myHttpServer.SessionInfo{Ctx: context.Background()}, true
}
