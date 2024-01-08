package useraccountsrv

import (
	"ibingli.com/internal/app/example/useraccount/useraccountdm"
	"ibingli.com/internal/pkg/myDatabase"
	"ibingli.com/internal/pkg/myHttpServer"
	"ibingli.com/internal/pkg/myLog"
)

type Repository struct {
	logger myLog.Iface
}

func newRepo(logger myLog.Iface) *Repository {
	return &Repository{
		logger: logger,
	}
}

func (repo *Repository) Create(ctx *myHttpServer.SessionInfo, db myDatabase.Iface, e *useraccountdm.UseraccountEntity) (insertedRecordId int64, err error) {
	return myDatabase.Create[useraccountdm.UseraccountEntity](db, repo.logger, e)
}

func (repo *Repository) PhysicalDeleteById(ctx *myHttpServer.SessionInfo, db myDatabase.Iface, id int64) error {
	return myDatabase.PhysicalDeleteById[useraccountdm.UseraccountEntity](db, repo.logger, id)
}

func (repo *Repository) LogicalDeleteById(ctx *myHttpServer.SessionInfo, db myDatabase.Iface, id int64) error {
	return myDatabase.LogicalDeleteById[useraccountdm.UseraccountEntity](ctx.Uid, db, repo.logger, id)
}

func (repo *Repository) UpdateById(ctx *myHttpServer.SessionInfo, db myDatabase.Iface, e *useraccountdm.UseraccountEntity) error {
	return myDatabase.UpdateById[useraccountdm.UseraccountEntity](db, repo.logger, e)
}

func (repo *Repository) SelectById(ctx *myHttpServer.SessionInfo, db myDatabase.Iface, id int64) (dst *useraccountdm.UseraccountEntity, err error) {
	return myDatabase.SelectById[useraccountdm.UseraccountEntity](db, repo.logger, id)
}

func (repo *Repository) SelectPage(ctx *myHttpServer.SessionInfo, db myDatabase.Iface, filter *myDatabase.Filter) ([]useraccountdm.UseraccountEntity, int64, error) {
	return myDatabase.SelectPage[useraccountdm.UseraccountEntity](db, repo.logger, filter)
}

func (repo *Repository) TransactionExample(ctx *myHttpServer.SessionInfo, db myDatabase.Iface) error {
	err := myDatabase.DoTransaction(db, func(tx myDatabase.Iface) error {

		return nil
	})

	return err
}
