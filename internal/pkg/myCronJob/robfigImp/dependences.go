// 用于声明本模块所需的接口, 具体实现将在运行时注入
package robfigImp

type loggerIface interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}
