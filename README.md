# comments模块

从bookinfo项目借鉴而来。 [bookinfo](https://github.com/GxlZ/bookinfo)
项目基于[go-kit](https://github.com/go-kit/kit)搭建,并集成常用组件.  

[myservices](https://github.com/FengGeSe/myservices)项目定义了微服务相关的基础设施。

# 功能list
---
- [x] 服务崩溃恢复
- [x] 基础代码自动生成
- [x] 依赖库管理
- [x] 实时编译
- [x] 命令行支持
- [x] grpc支持
- [x] grpc => http 协议转换支持
- [x] 服务熔断
- [x] 访问频率限制
- [x] prometheus采集支持
- [x] docker构建
- [x] 自定义中间件
- [x] 实时debug图表信息
- [x] pprof分析器，图表化
- [x] 生成火焰图 
- [x] 服务优雅退出 graceful
- [x] yaml配置文件支持
- [x] env配置文件支持
- [x] zipkin全链路追踪
- [x] test demo
- [x] benchmark demo
- [x] 多环境构建

# 文件目录
```
.
├── Dockerfile //docker构建文件
├── README.md 
├── cmd //服务启动&访问入口
|   ├── servername //客户端
|   └── servername-server //服务端
├── conf //多环境配置文件保存目录
|   ├── k8s.yaml //k8s模式配置文件
|   ├── container.yaml //容器模式下配置文件
|   └── local.yaml //本地开发模式下配置文件
├── global //全局生效 服务实例&变量&配置
|   ├── conf.go //实例化配置文件
|   ├── db.go //数据库相关
|   ├── errors.go //错误信息配置
|   ├── global.go //全局变量&配置 入口
|   ├── grpcclient.go //grpc客户端
|   ├── logger.go //日志实例
|   ├── pid.go //记录运行时服务pid
|   ├── redis.go //redis相关
|   └── zipkin.go //zipkin相关
├── handlers //业务逻辑目录
|   ├── handlers.go //最终服务实现入口文件
|   ├── hooks.go //钩子(graceful在此实现)
|   ├── middlewares.go //中间件加载文件
|   ├── mw_ep_name... //endpoint类型中间件
|   └── mw_svc_name... //svc类型中间件
├── lib //工具目录
|   ├── pid.go //生成pid
├── models //数据模型
|   ├── migrate.go //数据库迁移文件
|   ├── modelname.go //模型文件
|   └── ... //模型文件
├── pb //pb文件保存目录
├── runtime //保存程序运行时数据
|   ├── pid //服务运行pid
|   ├── logs //日志保存目录
|   └── ... //其他运行时数据
├── svc
|   ├── client //服务访问相关
|   |   ├── cli //命令行方式
|   |   ├── grpc //grpc方式
|   |   └── http //http方式
|   ├── server //服务启动相关
|   |   └── run.go //服务启动逻辑
|   ├── endpoints.go //endpoints实现
|   ├── transport_grpc.go //transport grpc实现
|   └── transport_http.go //transport http实现
├── vendor   // 第三方依赖
├── .realize.yaml //实时自动编译配置文件
├── .gitignore //git忽略规则文件
└── .env //项目运行环境变量
```
