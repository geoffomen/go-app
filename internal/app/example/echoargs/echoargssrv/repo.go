package echoargssrv

import (
	"ibingli.com/internal/app/example/echoargs/echoargsdm"
	"ibingli.com/internal/pkg/myDatabase"
	"ibingli.com/internal/pkg/myHttpServer"
	"ibingli.com/internal/pkg/myLog"
)

type EchoargsRepository struct {
	logger myLog.Iface
}

func newRepo(logger myLog.Iface) *EchoargsRepository {
	return &EchoargsRepository{
		logger: logger,
	}
}

func (repo *EchoargsRepository) Create(ctx *myHttpServer.SessionInfo, db myDatabase.Iface, e *echoargsdm.EchoargsEntity) (insertedRecordId int64, err error) {
	return myDatabase.Create[echoargsdm.EchoargsEntity](db, repo.logger, e)
}

func (repo *EchoargsRepository) PhysicalDeleteById(ctx *myHttpServer.SessionInfo, db myDatabase.Iface, id int64) error {
	return myDatabase.PhysicalDeleteById[echoargsdm.EchoargsEntity](db, repo.logger, id)
}

func (repo *EchoargsRepository) LogicalDeleteById(ctx *myHttpServer.SessionInfo, db myDatabase.Iface, id int64) error {
	return myDatabase.LogicalDeleteById[echoargsdm.EchoargsEntity](ctx.Uid, db, repo.logger, id)
}

func (repo *EchoargsRepository) UpdateById(ctx *myHttpServer.SessionInfo, db myDatabase.Iface, e *echoargsdm.EchoargsEntity) error {
	return myDatabase.UpdateById[echoargsdm.EchoargsEntity](db, repo.logger, e)
}

func (repo *EchoargsRepository) SelectById(ctx *myHttpServer.SessionInfo, db myDatabase.Iface, id int64) (dst *echoargsdm.EchoargsEntity, err error) {
	return myDatabase.SelectById[echoargsdm.EchoargsEntity](db, repo.logger, id)
}

func (repo *EchoargsRepository) SelectPage(ctx *myHttpServer.SessionInfo, db myDatabase.Iface, filter *myDatabase.Filter) ([]echoargsdm.EchoargsEntity, int64, error) {
	return myDatabase.SelectPage[echoargsdm.EchoargsEntity](db, repo.logger, filter)
}

func (repo *EchoargsRepository) TransactionExample(ctx *myHttpServer.SessionInfo, db myDatabase.Iface) error {
	err := myDatabase.DoTransaction(db, func(tx myDatabase.Iface) error {

		return nil
	})

	return err
}
