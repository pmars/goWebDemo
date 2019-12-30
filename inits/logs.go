package inits

import (
	"fmt"
	"strings"
	"time"

	"github.com/pmars/gotools"

	"goWebDemo/conf"
	"goWebDemo/tools"

	"github.com/pmars/beego/logs"
	"github.com/pmars/gotools/log"
	"github.com/pmars/gotools/wechat"
)

var OutboundIP string

func initLog() {
	logs.Async()
	logs.SetLogFuncCall(true)
	logs.SetLevel(logs.LevelDebug)
	logs.SetLogFuncCallDepth(4)
	logs.SetLogger(logs.AdapterConsole)
	logs.SetLogger(logs.AdapterFile, `{"filename":"./logs/pkgame.log","daily":true}`)
	logs.SetLogger(tools.AdapterNsqPush) // 启用nsq收集日志

	initWechatErrLog()

	OutboundIP = gotools.GetOutboundIP().String()
}

func initPushData(msg string) map[string]*wechat.TemplateInfo {
	parts := strings.SplitN(msg, " ", 3)
	return map[string]*wechat.TemplateInfo{
		"first":    {Value: "读伴儿分级阅读运行错误告警", Color: "#173177"},
		"keyword1": {Value: OutboundIP, Color: "#173177"},
		"keyword2": {Value: time.Now().Format("2006-01-02 15:04:05"), Color: "#173177"},
		"keyword3": {Value: parts[1], Color: "#173177"},
		"keyword4": {Value: parts[2], Color: "#173177"},
		"remark":   {Value: "服务器运行状态监控消息，请持续关注", Color: "#173177"},
	}
}

func initWechatErrLog() {
	if conf.Config.WechatPush.Need {
		jsonStr := fmt.Sprintf(`{
			"appId":"%v",
			"secret":"%v",
			"level":3,
			"tmpId":"%v",
			"wxIds":%v,
			"redisConn":"%v",
			"redisAuth":"%v",
			"redisKey":"%v"
		}`, conf.Config.WechatPush.AppId, conf.Config.WechatPush.Secret, conf.Config.WechatPush.TmpId,
			conf.Config.WechatPush.UserIds, conf.Config.WechatPush.RedisConn,
			conf.Config.WechatPush.RedisAuth, conf.Config.WechatPush.RedisKey)
		if err := log.InitWechatPush(jsonStr, initPushData); err != nil {
			logs.Error("initWechatErrLog ERROR:%v", err)
		}
		logs.Debug("Init Wechat ERROR Log PUSH Service Done!!!")
	}
}
