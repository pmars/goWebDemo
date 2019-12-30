package tools

import (
	"github.com/gin-gonic/gin"
)

const (
	ReturnMapKey = "returnMap"
)

// 设置返回内容的错误信息，返回内容之前调用
func (rm *ResultMap) SetReturnMsg(c *gin.Context) {
	errCode := GetCode(c)

	rm.Code = errCode.Code
	rm.Msg = errCode.CnMsg
	rm.MsgEn = errCode.EnMsg

	if rm.Code == Success.Code {
		requestId, _ := c.Get(RequestKey)
		rm.Result["request_id"] = requestId
	}
}

// 获取http.context中保存的Return信息
// 接口主要在Middleware和Controller中使用，获取对应的信息，返回信息
func GetResultMap(c *gin.Context) *ResultMap {
	retI, exists := c.Get(ReturnMapKey)
	if !exists {
		return setReturnMap(c)
	}
	ret, ok := retI.(*ResultMap)
	if !ok {
		return setReturnMap(c)
	}
	return ret
}

// 设置Return信息到http.context中
func setReturnMap(c *gin.Context) *ResultMap {
	ret := &ResultMap{Code: Success.Code}
	ret.Result = make(map[string]interface{})
	c.Set(ReturnMapKey, ret)
	return ret
}

type ResultMap struct {
	Code   int                    `json:"code"`
	Msg    string                 `json:"msg"`
	MsgEn  string                 `json:"msg_en"`
	Result map[string]interface{} `json:"result"`
}
