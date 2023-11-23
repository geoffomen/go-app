// 声明本模块所依赖的接口规范，这些接口规范由其它模块实现，并由main模块注入到当前模块中。
package echoargssrv

import (
	"example.com/internal/app/common/base/vo"
	"example.com/internal/app/example/useraccount/useraccountdm"
)

type useraccountSrvIface interface {
	Register(ctx vo.SessionInfo, args useraccountdm.CreateRequestDto) (int64, error)
}
