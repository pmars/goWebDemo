package tools

import "github.com/gin-gonic/gin"

const (
	CodeKey        = "returnCode"
	RequestKey     = "requestKey"
	UserKey        = "userObj"
	AdapterNsqPush = "nsq_push"
)

var (
	Success            = &ReturnCode{0, "Success", "success"}
	StatusUnauthorized = &ReturnCode{401, "权限错误", "StatusUnauthorized"}

	ErrArgs      = &ReturnCode{10000, "参数错误", "args error"}
	ErrVersion   = &ReturnCode{10001, "版本错误", "version error"}
	ErrServer    = &ReturnCode{50000, "网络开小差了，一会儿再来试试吧~", "server error"}
	ErrServerSys = &ReturnCode{50001, "系统出现了个小臭虫，工程师们正在除虫，一会儿再来试试吧~", "server error"}

	ErrRewardFull       = &ReturnCode{60101, "打卡奖励次数已经用完", "the number of rewards is full"}
	ErrRewardNotArrived = &ReturnCode{60102, "打卡奖励时间未到", "reward time has not arrived"}

	ErrAlreadySignin = &ReturnCode{60201, "今天已经签到过", "already signin"}

	ErrRoomId = &ReturnCode{60301, "游戏房间ID错误", "room id error"}
)

// 此方法主要在Controller中使用，设置错误码专用
func SetCode(c *gin.Context, code *ReturnCode) {
	c.Set(CodeKey, code)
}

// 在返回结果之前调用，获取错误码，如果没有设置，默认成功
func GetCode(c *gin.Context) *ReturnCode {
	code, exist := c.Get(CodeKey)
	if !exist {
		return Success
	}
	return code.(*ReturnCode)
}

type ReturnCode struct {
	Code  int    `json:"code"`
	CnMsg string `json:"msg"`
	EnMsg string `json:"msg_en"`
}

func (code *ReturnCode) SetMsg(cnMsg, enMsg string) {
	code.CnMsg = cnMsg
	code.EnMsg = enMsg
}
