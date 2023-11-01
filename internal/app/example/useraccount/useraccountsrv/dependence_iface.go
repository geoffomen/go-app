// 声明本模块所依赖的接口规范，这些接口规范由其它模块实现，并由main模块注入到当前模块中。
package useraccountsrv

import (
	"example.com/internal/app/common/base/vo"
)

type useraccountRepo interface {
	Create(ctx vo.SessionInfo, e UseraccountEntity) (insertedRecordId int64, err error)
	LogicalDeleteById(ctx vo.SessionInfo, id int64) error
	UpdateById(ctx vo.SessionInfo, e UseraccountEntity) error
	SelectById(ctx vo.SessionInfo, id int64) (dst *UseraccountEntity, err error)
	SelectPage(ctx vo.SessionInfo, condetions []string, orderBy string, offset int64, limit int64) ([]UseraccountEntity, int64,  error)
}
