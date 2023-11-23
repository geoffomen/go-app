// 声明本模块所依赖的接口规范，这些接口规范由其它模块实现，并由main模块注入到当前模块中。
package useraccountsrv

import (
	"example.com/internal/app/common/base/vo"
	"example.com/internal/app/example/echoargs/echoargsdm"
	"example.com/internal/app/example/useraccount/useraccountdm"
)

type useraccountRepo interface {
	Create(ctx vo.SessionInfo, e useraccountdm.UseraccountEntity) (insertedRecordId int64, err error)
	LogicalDeleteById(ctx vo.SessionInfo, id int64) error
	UpdateById(ctx vo.SessionInfo, e useraccountdm.UseraccountEntity) error
	SelectById(ctx vo.SessionInfo, id int64) (dst *useraccountdm.UseraccountEntity, err error)
	SelectPage(ctx vo.SessionInfo, condetions []string, orderBy string, offset int64, limit int64) ([]useraccountdm.UseraccountEntity, int64, error)
}

type ecoArgsSrvIface interface {
	Echo(ctx vo.SessionInfo, args echoargsdm.EhcoReqDto) (*echoargsdm.EchoRspDto, error)
}
