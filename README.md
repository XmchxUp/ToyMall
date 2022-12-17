# ToyMall


# 项目的主要依赖：
Golang V1.16
- gin
- gorm
- mysql
- redis
- ini
- jwt-go
- crypto
- logrus
- qiniu-go-sdk
- dbresolver

# 项目结构
```
gin-mall/
├── api
├── cache
├── conf
├── dao
├── doc
├── middleware
├── model
├── pkg
│  ├── e
│  └── util
├── routes
├── serializer
└── service
```
- api : 用于定义接口函数
- cache : 放置redis缓存
- conf : 用于存储配置文件
- dao : 对持久层进行操作
- doc : 存放接口文档
- middleware : 应用中间件
- model : 应用数据库模型
- pkg/e : 封装错误码
- pkg/util : 工具函数
- routes : 路由逻辑处理
- serializer : 将数据序列化为 json 的函数
- service : 接口函数的实现