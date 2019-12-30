package models

import (
	"goWebDemo/dao"
	"goWebDemo/log"

	"github.com/pmars/gotools"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

var (
	SQLGetUser         = "select * from pg_user where id = ?"
	SQLGetUserByOpenId = "select * from pg_user where open_id = ?"
	SQLAddUser         = "insert into pg_user (open_id) values (?)"
	SQLAddUserLogin    = "insert into pg_user_login (user_id, session_key, device) values (?, ?, ?)"
	SQLAddUserOpen     = "insert into pg_user_open (user_id) values (?)"
	SQLUpdateUserInfo  = "update pg_user set nick_name=?,avatar_url=?,country=?,province=?,city=?,language=?,gender=? where id=?"
)

type User struct {
	Id        int64
	OpenId    string `json:"openId"`    // 小程序对应的OpenId
	NickName  string `json:"nickName"`  // 用户名称
	AvatarUrl string `json:"avatarUrl"` // 用户头像
	Country   string `json:"country"`   //
	Province  string `json:"province"`  //
	City      string `json:"city"`      //
	Language  string `json:"language"`
	Gender    int    `json:"gender"`
}

func MGetUserByOpenId(c *gin.Context, openId string) *User {
	var user User
	if exist, err := dao.Engine.SQL(SQLGetUserByOpenId, openId).Get(&user); err != nil || !exist {
		log.Debug(c, "MGetUserByOpenId openId:%v exist:%v error:%v", openId, exist, err)
		return nil
	}
	return &user
}

func MGetUserById(c *gin.Context, userId int64) *User {
	var user User
	if exist, err := dao.Engine.SQL(SQLGetUser, userId).Get(&user); err != nil || !exist {
		log.Debug(c, "MGetUser userId:%v exist:%v error:%v", userId, exist, err)
		return nil
	}
	return &user
}

func MAddUser(c *gin.Context, session *xorm.Session, userP *User) (err error) {
	if result, err := session.Exec(SQLAddUser, userP.OpenId); err != nil {
		log.Error(c, "MAddUser OpenId:%v error:%v", userP.OpenId, err)
	} else {
		userP.Id, err = result.LastInsertId()
		if err != nil {
			log.Error(c, "MAddUser LastInsertId error:%v", err)
		}
	}
	return err
}

func MUpdateUserInfo(c *gin.Context, session *xorm.Session, userP *User) error {
	if _, err := session.Exec(SQLUpdateUserInfo, userP.NickName, userP.AvatarUrl, userP.Country, userP.Province, userP.City, userP.Language, userP.Gender, userP.Id); err != nil {
		log.Error(c, "MUpdateUserInfo user:%v error:%v", gotools.Data2Str(userP), err)
		return err
	}
	return nil
}

func MAddUserLogin(c *gin.Context, session *xorm.Session, userId int64, sessionKey, device string) {
	if _, err := session.Exec(SQLAddUserLogin, userId, sessionKey, device); err != nil {
		log.Error(c, "MAddUserLogin userId:%v sessionKey:%v device:%v error:%v", userId, sessionKey, device, err)
	} else {
		log.Debug(c, "MAddUserLogin userId:%v sessionKey:%v device:%v Success", userId, sessionKey, device)
	}
}

func MAddUserOpen(c *gin.Context, session *xorm.Session, userId int64) {
	if _, err := session.Exec(SQLAddUserOpen, userId); err != nil {
		log.Error(c, "MAddUserOpen userId:%v error:%v", userId, err)
	} else {
		log.Debug(c, "MAddUserOpen userId:%v Success", userId)
	}
}
