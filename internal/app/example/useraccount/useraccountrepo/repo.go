package useraccountrepo

import (
	"database/sql"

	"example.com/internal/app/common/base/entity"
	"example.com/internal/app/common/base/vo"
	"example.com/internal/app/example/useraccount/useraccountdm"
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

func (repo *UseraccountRepo) Create(ctx vo.SessionInfo, e useraccountdm.UseraccountEntity) (insertedRecordId int64, err error) {
	return entity.Create[useraccountdm.UseraccountEntity](ctx, repo.db, repo.logger, e)
}

func (repo *UseraccountRepo) PhysicalDeleteById(ctx vo.SessionInfo, id int64) error {
	return entity.PhysicalDeleteById[useraccountdm.UseraccountEntity](ctx, repo.db, repo.logger, id)
}

func (repo *UseraccountRepo) LogicalDeleteById(ctx vo.SessionInfo, id int64) error {
	return entity.LogicalDeleteById[useraccountdm.UseraccountEntity](ctx, repo.db, repo.logger, id)
}

func (repo *UseraccountRepo) UpdateById(ctx vo.SessionInfo, e useraccountdm.UseraccountEntity) error {
	return entity.UpdateById[useraccountdm.UseraccountEntity](ctx, repo.db, repo.logger, e)
}

func (repo *UseraccountRepo) SelectById(ctx vo.SessionInfo, id int64) (dst *useraccountdm.UseraccountEntity, err error) {
	return entity.SelectById[useraccountdm.UseraccountEntity](ctx, repo.db, repo.logger, id)
}

func (repo *UseraccountRepo) SelectPage(ctx vo.SessionInfo, conditions []string, orderBy string, offset int64, limit int64) ([]useraccountdm.UseraccountEntity, int64, error) {
	return entity.SelectPage[useraccountdm.UseraccountEntity](ctx, repo.db, repo.logger, conditions, orderBy, offset, limit)
}
