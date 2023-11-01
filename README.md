# 架构目标
### 软件项目的腐化过程
- 产品方忽略架构的价值，要求软件产品尽快工作起来

    产品方认为软件正常工作最重要，而且他们所提出的一系列的变更需求的范畴都是类似的，因此成本也应该是固定的，从而给了不符合实际的时间要求到研发。但是从研发者角度来看，系统用户持续不断的变更需求就像是要求他们不停地用一堆不同形状的拼图块，拼成一个新的形状。整个拼图的过程越来越困难，因为现有系统的形状永远和需求的形状不一致。研发缺乏时间认真设计架构，牺牲架构未来的灵活度来快速实现当前功能，导致系统越来越难以维护，最终无法修改。
- 工程师过度自信，持续低估那些好的、良好设计的、整洁的代码的重要性

    他们普遍用一句话来欺骗自己：“我们可以未来再重构代码，产品上线最重要！”但是结果大家都知道，产品上线以后重构工作就再没人提起。市场的压力永远也不会消退，作为先上市的产品，后面有无数的竞争对手追赶，必须要比他们跑得更快才能保持领先。所以，重构的时机永远不会再有。工程师们忙于完成新功能，新功能做不完，哪有时间重构老的代码？循环往复，系统成了一团乱麻，代码依赖关系错综复杂，组件相互耦合紧密，导致不管多么小的改动都需要数周的恶战才能完成，生产效率持续急速下降，直至为零。
- 软件系统的设计如此之差，系统中到处充满了腐朽的设计和连篇累牍的恶心代码，处处都是障碍，让整个团队的士气低落，用户天天痛苦，经理们手足无措。
- 软件系统因其架构腐朽不堪，而导致团队流失，部门解散，甚至公司倒闭。
### 目标
创建出一个可以**让功能实现起来更容易、修改起来更简单、扩展起来更轻松的软件架构**，用最小的人力成本来满足构建和维护该系统的需求。

### 设计原则
本架构是对《架构整洁之道》书中提出的精神的粗陋理解和粗陋实现，利用多态技术和依赖倒置原则管理代码的依赖关系。

在代码的执行流程上，一个组件想要调用另一个组件的函数，必须知道被调用函数的具体信息，这就产生了强耦合。强耦合带来的坏处是，首先不能轻松替换被调用者，其次被调用者的改动须要考虑对调用者产生的影响。但是在代码组织上，我们可以通过多态技术，将依赖关系反转过来，让被调用者依赖调用者的接口规范，从而打破强耦合。具体实现就是，调用者声明自己想要调用的函数的接口规范，然后由被调用者实现该接口规范，最后在执行过程中将被调用者注入到调用者中，这就是多态技术。代码的依赖关系与执行时依赖关系倒置了，这就是依赖倒置原则。

### 特色
- 提供模块代码生成功能
- 提供参数绑定功能
- 提供数据库的CURD基本操作
- 优先选用编程语言提供的功能，而非第三方实现
- 一键编译软件
- 通过docker一键运行项目

# 体验例子代码
- 在项目根目录下执行` sh build/package/example/build.sh`, 会在build/temp/目录生成一个可执行文件**example.o**
- 运行生成的可执行文件`build/temp/example.o`，完成后即可使用rest client(如postman)调用服务端的api接口。
- 可以运行`go test test/example_api_test.go`执行简单API测试。

# 快速开始编写业务
- 在项目根目录执行`go run tools/gen_module_files.go`
- 根据提示，输入应用名称，例如：examples，按回车确认
- 根据提示，输入模块名称，例如：test, 按回车确认
- 程序会在`internal/app/{应用名称}/{模块名称}`生成模块相关文件。可按需修改`tools/gen_module_files.go`和`tools/moduletemplates/`下的文件来控制生成的文件
- 在`internal/app/{应用名称}/{模块名称}/{模块名称}srv/service_base.go`文件接收所依赖的接口的具体实现的注入
- 在`internal/app/{应用名称}/{模块名称}/{模块名称}srv/service.go`文件实现业务逻辑
- 在`internal/app/{应用名称}/{模块名称}/{模块名称}srv/entity.go`文件定义业务实体
- 在`internal/app/{应用名称}/{模块名称}/{模块名称}srv/req_dto.go`文件定义请求DTO
- 在`internal/app/{应用名称}/{模块名称}/{模块名称}srv/rsp_dto.go`文件定义响应DTO
- 在`internal/app/{应用名称}/{模块名称}/{模块名称}srv/vo.go`文件定义值对象
- 在`internal/app/{应用名称}/{模块名称}/{模块名称}srv/dependence_ifce.go`文件按需定义依赖的接口以及一些常量
- 在`internal/app/{应用名称}/{模块名称}/{模块名称}repo/repo.go`文件定义数据的存储操作
- 在`internal/app/{应用名称}/{模块名称}/{模块名称}repo/dependence_iface.go`文件按需定义数据的存储操作所依赖的接口以及一些常量
- 在`internal/app/{应用名称}/{模块名称}/{模块名称}ad/{模块名称}http/controller_base.go`文件注入依赖
- 在`internal/app/{应用名称}/{模块名称}/{模块名称}ad/{模块名称}http/controller.go`文件定义路由
- 在`cmd/{应用名称}/main.go`文件向协议处理框架(如http)注册路由
- 完成。可参考`internal/app/examples/`下面的例子
- 如果单文件下代码量多，建议拆分成多个逻辑相关的小文件，方便导航。
- 按照面向对象的封装原则, 各模块之间的交互应当仅通过接口定义、依赖注入的方式进行。

# 编译运行
有两种方式，建议使用容器的方式。
- 使用容器一站式启动, 所需的依赖都会启动，如：数据库和nginx
    - 安装docker、docker-compose-plugin, make
    - 在项目根目录下执行`sh build/container/{应用名称}/build.sh`打包容器，根据脚本提示使用容器。
- 编译为可执行文件, 须要自己预先启动所须依赖
    - 编译：在项目根目录执行`sh build/package/{应用名称}/build.sh`
    - 运行：在项目根目录执行`build/temp/{应用名称}.o`或指定配置文件`build/temp/{应用名称}.o --profile=dev` （忽略配置文件后缀，会从当前目录和当前目录下的configs目录查找后缀为.yml的同名配置文件）
    - 可按需修改`Makefile`和`scripts/`下的脚本


# 代码组织结构
```
├── assets  # 项目运行时所需资源，如：字体、excel模板
│   └── README.md
├── build  # 项目构建相关
│   ├── container  # 容器
│   │   └── example
│   │       └── Dockerfile
│   ├── package  # 打包
│   │   └── example
│   ├── README.md
│   └── temp
├── cmd  # 程序入口
│   ├── example
│   │   └── main.go
│   └── README.md
├── configs  # 配置文件
├── deployments  # 部署相关
│   ├── container
│   ├── mysql
│   └── nginx
├── docs  # 文档
│   ├── git flow分支策略.md
│   └── README.md
├── internal  
│   ├── app  # 子应用目录
│   │   ├── common  # 通用应用
│   │   │   └── base  # 通用基础功能模块
│   │   │       ├── entity
│   │   │       └── vo
│   │   └── example  # 例子应用
│   │       ├── echoargs  # 应用功能模块
│   │       │   ├── echoargsctl
│   │       │   ├── echoargssrv
│   │       │   └── README.md
│   │       └── useraccount  # 应用功能模块
│   │           ├── README.md
│   │           ├── useraccountctl  # 
│   │           ├── useraccountrepo  # 数据操作
│   │           └── useraccountsrv  # 业务逻辑
│   ├── pkg  # 实用工具
│   └── README.md
├── pkg
│   └── README.md
├── scripts  # 方便脚本
├── test  # 额外的测试文件
│   ├── example_api_test.go
│   ├── example_curl_test.md
│   └── README.md
├── tools  # 方便工具
│   ├── gen_module_files.go  # 控制模块文件生成
│   ├── moduletemplates  # 模块模板文件
│   └── README.md
├── vendor  # 第三方依赖
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
├── README.md
└── website  # 静态网页文件
    └── example

```
