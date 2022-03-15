# 特性
本项目是按功能分包的模块化设计的API服务框架，封装了常用的功能并遵循依赖倒置原则，致力于进行快速的业务研发。

- 支持参数绑定，采取控制反转的方式注入依赖，方便进行单元测试；
- 提供了常见组件的抽象层，支持灵活地替换底层实现，避免了日后依赖更改时的散弹式修改；
- 支持 rate 接口限流
- 支持 cors 接口跨域
- 支持 jwt 接口鉴权
- 支持 RESTful API 返回值规范
- 支持 zap 日志收集
- 支持 viper 配置文件解析
- 支持 cron 定时任务
- 支持 gorm 数据库组件
- 支持功能模块框架生成


# 体验例子代码
- 在examples目录下执行`go build -mod=vendor -o main.o main.go`, 会在当前工作目录生成一个可执行文件**main.o**
- 运行生成的可执行文件`./main.o`，完成后即可使用rest client(如postman)调用服务端的api接口。


# 快速开始编写业务
- 在项目根目录执行`go run tools/gen_module_files.go`
- 根据提示，输入模块名称，按回车确认
- 程序会在`internal/app/myapp/`生成模块相关文件。可按需修改`tools/gen_module_files.go`和`tools/moduletemplates/`下的文件来控制生成的文件
- 在`internal/app/myapp/<模块名称>/*imp/service.go`文件实现业务逻辑
- 在`internal/app/myapp/<模块名称>/req_dto.go`文件定义请求DTO
- 在`internal/app/myapp/<模块名称>/rsp_dto.go`文件定义响应DTO
- 在`internal/app/myapp/<模块名称>/*ctl/controller.go`文件定义路由
- 在`internal/app/myapp/<模块名称>/iface.go`文件按需定义模块对外暴露的函数
- 在`cmd/myapp/main.go`文件向协议处理框架(如http)注册路由
- 完成。可参考`examples/`下面的例子
- 如果单文件下代码量多，建议将文件拆分成多个小文件，方便导航，参考`examples/hello`模块


# 编译运行
有两种方式，建议使用容器的方式。
- 使用容器一站式启动
    - 安装podman、make
    - 在项目根目录下执行`make myapp-container-run`启动容器。
    - 在项目根目录下执行`make myapp-db-container-data-import`导入数据(可按需修改`Makefile`里的数据文件路径)。
    - 在项目根目录下执行`make myapp-pod-clean`可删除正在运行的容器(注意，数据库的数据也会删除，如果想保存可执行`make myapp-db-container-data-export`，下次再重新导入)。
- 编译为可执行文件
    - 编译：在项目根目录执行`make myapp`
    - 运行：在项目根目录执行`build/temp/myapp.o`或指定配置文件`build/temp/myapp.o --profile=dev` （忽略配置文件后缀，会从当前目录和当前目录下的configs目录查找后缀为.yml的同名配置文件）
    - 可按需修改`Makefile`和`scripts/`下的脚本


# 代码组织结构
```
├── build (构建时的工作目录)
│   └── temp
├── cmd (程序入口目录)
│   ├── myapp
│   │   └── main.go (入口函数)
│   └── README.md
├── configs (配置文件目录)
│   ├── dev.yml
│   └── README.md
├── docs (项目文档目录)
│   └── README.md
├── examples (项目例子)
│   ├── useracc
│   │   ├── useraccctl (路由)
│   │   │   └── controller.go
│   │   ├── useraccimp (实现)
│   │   │   ├── entity.go
│   │   │   └── service.go
│   │   ├── iface.go (模块接口)
│   │   ├── request_dto.go (请求DTO)
│   │   └── response_dto.go (响应DTO)
│   ├── user
│   ├── hello
│   ├── README.md
│   ├── example.yml
│   ├── main.go
│   └── api_test.go
├── internal
│   ├── app (业务代码)
│   │   └── myapp
│   └── pkg (对内基础组件)
├── pkg (对外基础组件)
│   ├── README.md
│   ├── config (配置读取封装)
│   │   ├── config_iface_test.yml
│   │   ├── iface.go (接口)
│   │   ├── iface_test.go
│   │   ├── mapimp (接口的某个实现)
│   │   │   └── service.go
│   │   └── viperimp (接口的某个实现)
│   │       └── service.go
│   ├── cronjob (定时任务封装)
│   ├── database (数据库封装)
│   ├── digestutil (加解密实用工具)
│   ├── httpclient (httpclient实用工具)
│   ├── mq (消息队列封装)
│   ├── myerr (错误封装)
│   ├── mylog (日志封装)
│   ├── stringutil (字符串实用工用)
│   └── httpfw (http协议处理框架)
│       ├── ginimp
│       └── iface.go
├── scripts (方便脚本)
│   ├── build.sh
│   ├── clean.sh
│   └── README.md
├── tools (辅助工具)
│   ├── gen_module_files.go
│   ├── moduletemplates (代码生成所用的模板)
│   │   ├── controller.tmpl
│   │   ├── iface.tmpl
│   │   ├── req_dto.tmpl
│   │   ├── rsp_dto.tmpl
│   │   └── service.tmpl
│   └── README.md
├── LICENSE
├── Makefile 
├── README.md (项目说明文档)
├── go.mod
├── go.sum
└── vendor (第三方依赖库)
```
