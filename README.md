### 目录结构
```
├─configs  配置文件
├─docs     文档集合
├─global    全局变量
├─internal  内部模块
│  ├─dao    数据访问层
│  ├─middleware HTTP中间件
│  ├─model      模型层
│  ├─routers    路由相关逻辑
│  └─service    项目核心业务逻辑
├─pkg       项目相关模块包
├─storage   项目生成的临时文件
└─third_party   第三方资源库：Swagger Ui
```
### 用到的模块
```

web框架    github.com/gin-gonic/gin
配置文件    github.com/spf13/viper
接口文档    github.com/swaggo/swag
操作数据库  gorm.io/gorm
           gorm.io/driver/mysql
log日志     github.com/natefinch/lumberjack
接口校验    github.com/go-playground/validator/v10
限流控制    github.com/juju/ratelimit v1.0.1
发送邮件    gopkg.in/gomail.v2
配置热更新   github.com/fsnotify/fsnotify

```