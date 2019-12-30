# API PROJECT goWebDemo

这个部分主要将总结的一些通用代码整理一下，如果开启新的项目的话，可以在这个基础上直接复用

### 各模块用途

| 文件夹 | 用途说明 |
|-------|---------|
| app | main.go主函数，解决单元测试启动时的错误问题 |
| conf | 配置文件及初始化配置文件的部分 |
| controller | 各种接口文件，主要处理各种参数的判断等 |
| dao | 数据库相关的操作在此目录 |
| inits | 单独抽离出来的初始化部分，负责各种内容的初始化 |
| log | 重写的日志打印模块，增加requestId打印 |
| logs | 项目的日志存放位置 |
| models | 数据相关功能，包括各种数据的增删改查操作函数 |
| router | 路由部分，中间件都可以写在这里 |
| scripts | 部署脚本存放位置 |
| service | 主要功能文件，接口逻辑实现都可以写在这里，避免controller冗长 |
| static | HTML静态文件存放位置，解决SLB的监控 |
| tools | 项目中用到的工具函数，不依赖于项目中的任何部分 |
| .gitlab-ci.yml | gitlab ci/cd 自动部署配置文件 |


### 代码规范

其他需要注意的点：

+ Controller的函数里面只做参数的判断，还有最后的赋值
+ 所有的逻辑放到service里面
+ 即使再简单的请求从Controller到models也需要在service层中转一下
+ model层只负责和数据库，缓存相关的数据交互
+ 所有更新或者添加的model层的函数都需要从service层拿到session来进行，统一在service层来创建session
+ 所有的日子打印都需要用 log.Debug log.Info log.Error 来进行
+ 所有的函数都需要传 c *gin.Context 作为第一个参数
+ 每个Controller中的接口，都需要打印日志 log.Info(c, "CDeleteUserWord Start ")后面需要接上所有的参数
+ 每个实现的接口，都需要在Controller里面写详细的实现方案，如有改动及时更新，例：CRecognizeScore，CBookReading

```go
   export GO111MODULE=on
   export GOPROXY=https://goproxy.io
```

GoLand 设置支持go mod

   ![image-20190316154114325](https://dubaner-reading.oss-cn-beijing.aliyuncs.com/resource/ideajietu.png)


未完待续……

### END

