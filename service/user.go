package service

import (
	"goWebDemo/dao"
	"goWebDemo/log"
	"goWebDemo/models"
	"goWebDemo/tools"

	"github.com/gin-gonic/gin"
)

// 用户登录，获取用户信息，每次登陆调用
func SUserLogin(c *gin.Context, code, device string) (*models.User, string, string) {
	// 通过Code获取用户的OpenId，SessionKey
	openId, sessionKey := GetOpenId(c, code)
	log.Debug(c, "SUserLogin code:%v openId:%v sessionKey:%v", code, openId, sessionKey)
	if len(openId) == 0 {
		tools.SetCode(c, tools.ErrServer)
		return nil, "", ""
	}

	// 通过OpenId获取用户数据
	userP := models.MGetUserByOpenId(c, openId)

	var err error
	session := dao.Engine.NewSession()
	_ = session.Begin()
	defer func() {
		if err != nil {
			_ = session.Rollback()
		} else {
			_ = session.Commit()
		}
		session.Close()
	}()

	// 如果用户不存在，则创建新用户
	if userP == nil {
		userP = &models.User{
			OpenId: openId,
		}

		// 添加用户
		if err = models.MAddUser(c, session, userP); err != nil {
			return nil, "", ""
		}
	}

	// 将用户的登陆信息写入到数据库
	go models.MAddUserLogin(c, session, userP.Id, sessionKey, device)
	// 将用户打开小程序的记录也写入到数据库
	go models.MAddUserOpen(c, session, userP.Id)

	return userP, openId, sessionKey
}

// 更新用户信息
func SUpdateUserInfo(c *gin.Context, userP *models.User) {
	var err error
	session := dao.Engine.NewSession()
	_ = session.Begin()
	defer func() {
		if err != nil {
			_ = session.Rollback()
		} else {
			_ = session.Commit()
		}
		session.Close()
	}()

	if models.MUpdateUserInfo(c, session, userP) != nil {
		tools.SetCode(c, tools.ErrServer)
	}
}

// 用户打开小程序，记录
func SUserOpen(c *gin.Context, userP *models.User) {
	var err error
	session := dao.Engine.NewSession()
	_ = session.Begin()
	defer func() {
		if err != nil {
			_ = session.Rollback()
		} else {
			_ = session.Commit()
		}
		session.Close()
	}()

	// 将用户打开小程序的记录也写入到数据库
	go models.MAddUserOpen(c, session, userP.Id)
}
