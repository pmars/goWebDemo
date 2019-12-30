package controller

import (
	"goWebDemo/log"
	"goWebDemo/models"
	"goWebDemo/service"
	"goWebDemo/tools"
	"net/http"

	"github.com/pmars/gotools"

	"github.com/gin-gonic/gin"
)

// 获取http.context中保存的Return信息
// 接口主要在Middleware和Controller中使用，获取对应的信息，返回信息
func getUser(c *gin.Context) *models.User {
	userI, exists := c.Get(tools.UserKey)
	if !exists || userI == nil {
		c.AbortWithStatusJSON(http.StatusOK, tools.StatusUnauthorized)
		return nil
	}
	userP, ok := userI.(*models.User)
	if !ok || userP == nil {
		c.AbortWithStatusJSON(http.StatusOK, tools.StatusUnauthorized)
		return nil
	}
	return userP
}

/*
	通过Code用户的OpenId，SessionKey

	小程序中，首先获取缓存的OpenID，如果没有，这进行wx.login操作
	调用此接口，服务器到微信服务器换取OpenID，SessionKey返回前端
	后续接口，header中，需要有OpenId

*/
func CUserLogin(c *gin.Context) {
	var params struct {
		Code   string `json:"code"` //
		Device string `json:"device"`
	}
	if err := c.Bind(&params); err != nil {
		log.Debug(c, "CUserLogin Bind ERROR:%v", err)
		tools.SetCode(c, tools.ErrArgs)
		return
	}
	log.Info(c, "CUserLogin Start params:%v", gotools.Data2Str(params))

	// 获取用户信息
	_, openId, sessionKey := service.SUserLogin(c, params.Code, params.Device)

	if tools.GetCode(c) == tools.Success {
		data := tools.GetResultMap(c)
		data.Result["open_id"] = openId
		data.Result["session_key"] = sessionKey
	}
}

/*
	通过OpenId获取用户信息
*/
func CUserInfo(c *gin.Context) {
	userP := getUser(c)
	log.Info(c, "CUserInfo Start userId:%v", userP.Id)

	service.SUserOpen(c, userP)

	data := tools.GetResultMap(c)
	data.Result["user"] = userP
}

/*
	上传用户信息数据
*/
func CUpdateUserInfo(c *gin.Context) {
	userP := getUser(c)
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		log.Debug(c, "CUpdateUserInfo Bind ERROR:%v", err)
		tools.SetCode(c, tools.ErrArgs)
		return
	}
	log.Info(c, "CUpdateUserInfo Start userId:%v user:%v", userP.Id, gotools.Data2Str(user))

	user.Id = userP.Id
	service.SUpdateUserInfo(c, &user)

	if tools.GetCode(c) == tools.Success {
		data := tools.GetResultMap(c)
		data.Result["user"] = user
	}
}
