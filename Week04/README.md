学习笔记

项目在bbs目录

### bbs 目录结构
```sh
bbs        应用根目录
├─api           API 协议定义目录，xxapi.proto protobuf 文件，以及生成的 go 文件
├─cmd           应用程序可执行文件目录，主要可执行文件的名称保持和应用程序一致
│  ├─bbs-interface  对外的 BFF 服务，接受来自用户的请求，比如暴露了 HTTP/gRPC 接口
│  ├─bbs-service    对内的微服务，仅接受来自内部其他服务或者网关的请求，比如暴露了gRPC 接口只对内服务
│  ├─bbs-admin      后台管理
│  ├─bbs-job        常驻任务，流式处理，消费kafka消息等
│  ├─bbs-task       定时任务，crontab
├─configs       配置文件、默认配置等
├─internal      私有应用程序和库代码，不能被外部引用，internal目录go做了编译处理
│  ├─biz            业务逻辑代码类似于DDD domain
│  ├─data           业务数据访问接口，类似 DDD Repository
│  ├─service        类似 DDD Application
├─pkg           应用提供的公共库，可由其他项目引用，可以按照功能分类
├─test          额外的外部测试应用程序和测试数据
```