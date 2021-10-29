# App.

# 代码结构
```
├── build (构建时的工作目录)
│   └── temp
├── cmd (程序入口函数)
│   ├── myapp
│   │   └── main.go
│   └── README.md
├── configs (配置文件目录)
│   ├── dev.yml
│   └── README.md
├── docs (项目文档目录)
│   └── README.md
├── examples (项目例子)
│   ├── account (功能模块)
│   │   ├── accountctl (路由)
│   │   │   └── controller.go
│   │   ├── accountimp (实现)
│   │   │   ├── entity.go
│   │   │   └── service.go
│   │   ├── iface.go (模块接口)
│   │   ├── request_dto.go (请求DTO)
│   │   └── response_dto.go (响应DTO)
│   ├── user (功能模块)
│   ├── hello (功能模块)
│   ├── README.md (说明)
│   ├── example.yml (配置文件)
│   ├── main.go (例子入口函数)
│   └── api_test.go 
├── internal
│   ├── app (业务代码，结构与examples一致)
│   │   └── myapp
│   ├── pkg (基础组件)
│   │   ├── config (配置读取封装)
│   │   │   ├── config_iface_test.yml
│   │   │   ├── iface.go
│   │   │   ├── iface_test.go
│   │   │   ├── mapimp
│   │   │   │   └── service.go
│   │   │   └── viperimp
│   │   │       └── service.go
│   │   ├── cronjob (定时任封装)
│   │   ├── database (数据库封装)
│   │   ├── digestutil (加解密实用工具)
│   │   ├── httpclient (httpclient实用工具)
│   │   ├── mq (消息队列封装)
│   │   ├── myerr (错误封装)
│   │   ├── mylog (日志封装)
│   │   ├── stringutil (字符串实用工用)
│   │   ├── vo (基础值对象)
│   │   │   ├── base_req.go
│   │   │   ├── base_rsp.go
│   │   │   └── common_vo.go
│   │   └── webfw (http框架)
│   │       ├── ginimp
│   │       └── iface.go
│   └── README.md
├── scripts (方便脚本)
│   ├── build.sh
│   ├── clean.sh
│   └── README.md
├── tools (辅助工具)
│   ├── gen_module_files.go
│   ├── moduletemplates
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