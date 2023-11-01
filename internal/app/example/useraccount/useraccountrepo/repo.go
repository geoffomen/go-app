package useraccountrepo

import (
	"database/sql"

	"example.com/internal/app/common/base/entity"
	"example.com/internal/app/common/base/vo"
	"example.com/internal/app/example/useraccount/useraccountsrv"
)

type UseraccountRepo struct {
	db     *sql.DB
	logger LoggerIface
}

func New(db *sql.DB, logger LoggerIface) *UseraccountRepo {
	return &UseraccountRepo{
		db:     db,
		logger: logger,
	}
}

func (repo *UseraccountRepo) Create(ctx vo.SessionInfo, e useraccountsrv.UseraccountEntity) (insertedRecordId int64, err error) {
	return entity.Create[useraccountsrv.UseraccountEntity](ctx, repo.db, repo.logger, e)
}

func (repo *UseraccountRepo) PhysicalDeleteById(ctx vo.SessionInfo, id int64) error {
	return entity.PhysicalDeleteById[useraccountsrv.UseraccountEntity](ctx, repo.db, repo.logger, id)
}

func (repo *UseraccountRepo) LogicalDeleteById(ctx vo.SessionInfo, id int64) error {
	return entity.LogicalDeleteById[useraccountsrv.UseraccountEntity](ctx, repo.db, repo.logger, id)
}

func (repo *UseraccountRepo) UpdateById(ctx vo.SessionInfo, e useraccountsrv.UseraccountEntity) error {
	return entity.UpdateById[useraccountsrv.UseraccountEntity](ctx, repo.db, repo.logger, e)
}

func (repo *UseraccountRepo) SelectById(ctx vo.SessionInfo, id int64) (dst *useraccountsrv.UseraccountEntity, err error) {
	return entity.SelectById[useraccountsrv.UseraccountEntity](ctx, repo.db, repo.logger, id)
}

func (repo *UseraccountRepo) SelectPage(ctx vo.SessionInfo, conditions []string, orderBy string, offset int64, limit int64) ([]useraccountsrv.UseraccountEntity, int64, error) {
	return entity.SelectPage[useraccountsrv.UseraccountEntity](ctx, repo.db, repo.logger, conditions, orderBy, offset, limit)
}
