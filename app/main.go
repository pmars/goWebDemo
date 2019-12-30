package main

import (
	"goWebDemo/conf"
	"goWebDemo/router"

	"goWebDemo/inits"

	"github.com/gin-gonic/gin"
)

func main() {
	inits.InitAll()

	// 初始化gin
	gin.SetMode(conf.Config.Service.Mode)
	engine := gin.New()

	// 加载路由
	router.Middleware(engine)
	router.ApiV1(engine)

	// 做一个HTML路由
	engine.Static("/static", conf.Config.Service.Html)

	// 启动服务
	_ = engine.Run(conf.Config.Service.Addr)
}
