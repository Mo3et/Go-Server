# Go HTTP Server

[Articles](https://learnku.com/articles/400871212345)


![逻辑](https://cdn.learnku.com/uploads/images/202001/27/6964/kZ1Hmwtb14.png)

[Go Web 开发目录结构说明](https://studygolang.com/articles/32333)

```
goproj_org
├── todo.go    # main 入口
├── go.mod
├── go.sum
├── config    # 配置文件存放
│   └── config.yaml
├── controller  # api 接口逻辑，http handler
│   ├── DemoController.go
│   └── TodoController.go
├── model      # model定义
│   └── Todo.go
├── router    # 路由入口
│   ├── middleware  # gin框架 自定义中间件，中间件一般用于路由中，故放到这里
│   │   └── jwt.go
│   └── router.go
├── service 
│   ├── ConfigService.go
│   ├── DBService.go
│   └── TodoService.go
└── util      # 公共函数部分，比如加密，时间处理等
 └── encrypt.go</pre>
```