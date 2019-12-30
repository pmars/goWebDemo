package service

import (
	"encoding/json"
	"fmt"
	"goWebDemo/conf"
	"goWebDemo/log"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOpenId(c *gin.Context, code string) (openId, sessionKey string) {
	// 通过Code获取用户的OpenId，SessionKey
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v",
		conf.Config.MiniAppId, conf.Config.MiniSecret, code)
	log.Debug(c, "GetOpenId url:%v", url)
	bytes := getRequest(c, url)
	log.Debug(c, "GetOpenId result:%v", string(bytes))

	var params struct {
		OpenId     string `json:"openid"`
		SessionKey string `json:"session_key"`
	}
	if err := json.Unmarshal(bytes, &params); err != nil {
		log.Error(c, "GetOpenId error:%v", err)
	}
	return params.OpenId, params.SessionKey
}

func getRequest(c *gin.Context, url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Error(c, "getRequest url:%v err:%v", url, err)
		return nil
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(c, "getRequest url:%v err:%v", url, err)
		return nil
	}
	return bytes
}
