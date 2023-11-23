package useraccountsrv

import (
	"testing"

	"example.com/internal/app/common/base/vo"
	"example.com/internal/app/example/useraccount/useraccountdm"
	"github.com/stretchr/testify/assert"
)

type UseraccountRepoMock struct {
}

func (r *UseraccountRepoMock) Create(ctx vo.SessionInfo, e useraccountdm.UseraccountEntity) (insertedRecordId int64, err error) {
	return 1, nil
}

func (r *UseraccountRepoMock) LogicalDeleteById(ctx vo.SessionInfo, id int64) error {
	return nil
}

func (r *UseraccountRepoMock) UpdateById(ctx vo.SessionInfo, e useraccountdm.UseraccountEntity) error {
	return nil
}

func (r *UseraccountRepoMock) SelectById(ctx vo.SessionInfo, id int64) (dst *useraccountdm.UseraccountEntity, err error) {
	return nil, nil
}

func (r *UseraccountRepoMock) SelectPage(ctx vo.SessionInfo, condetions []string, orderBy string, offset int64, limit int64) ([]useraccountdm.UseraccountEntity, int64, error) {
	return nil, 0, nil
}

func TestService(t *testing.T) {
	repoMock := UseraccountRepoMock{}
	service := Service{
		repo: &repoMock,
	}

	args := useraccountdm.CreateRequestDto{
		Account:  "account1",
		Password: "123456",
	}
	id, err := service.Register(vo.SessionInfo{}, args)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, 0, id)
}
