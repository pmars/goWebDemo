package router

import (
	"goWebDemo/controller"

	"goWebDemo/inits"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func Middleware(engine *gin.Engine) {
	// 增加中间件
	engine.Use(gin.Recovery())
	engine.Use(cors)
	engine.Use(passMethods)
	engine.Use(setRequestId)
	engine.Use(defaultLog)
	engine.Use(outputHeader)
	engine.Use(checkEncrypt)
	engine.Use(checkOpenId)
	engine.Use(setReturn)
	engine.Use(favicon.New(*inits.LogoPath))
}

func ApiV1(engine *gin.Engine) {
	// 初始化路由信息
	apiV1 := engine.Group("/api/v1")
	{
		// 用户相关接口
		user := apiV1.Group("/user")
		{
			user.POST("/login", controller.CUserLogin)       // 用户登录接口
			user.POST("/info", controller.CUserInfo)         // 获取用户信息接口
			user.POST("/update", controller.CUpdateUserInfo) // 更新用户信息接口
		}

	}
}
