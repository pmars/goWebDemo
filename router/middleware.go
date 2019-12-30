package router

import (
	"goWebDemo/conf"
	"goWebDemo/log"
	"goWebDemo/models"
	"goWebDemo/tools"

	"bytes"
	"encoding/json"

	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"time"

	"github.com/pmars/gotools"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

// 过滤所有的Options, Head Method Requests
func passMethods(c *gin.Context) {
	if c.Request.Method == "OPTIONS" || c.Request.Method == "HEAD" {
		c.AbortWithStatus(200)
		return
	}
}

// 设置接口的RequestId
func setRequestId(c *gin.Context) {
	requestId := uuid.NewV4()
	c.Set(tools.RequestKey, strings.Replace(requestId.String(), "-", "", -1))
}

// 允许服务器跨域访问
func cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "OPTIONS,HEAD,GET,POST,DELETE,PUT")
	c.Header("Access-Control-Allow-Headers", "Authorization,Content-Type,Content-Range,Content-Disposition,token,device")
}

// 通用响应处理中间件
func setReturn(c *gin.Context) {
	// 调用请求
	c.Next()

	// 如果设置 notNeedSetReturn 参数为true，则不再执行后续流程（因为微信回调返回xml格式，用于绕过这里的包装）
	if c.GetBool("notNeedSetReturn") {
		return
	}
	// 获取Controller层设置的Return信息
	result := tools.GetResultMap(c)
	result.SetReturnMsg(c)
	log.Debug(c, "Middle Return: %v", gotools.Data2Str(result))
	c.JSON(200, result)
}

// 以下 URL Path 不需要检查 TOKEN
var noCheckTokenMap = map[string]bool{
	"/api/v1/user/login": true, // 用户登录
}

// 校验 OpenId 的中间件
func checkOpenId(c *gin.Context) {
	// 检查是否需要OpenId效验
	if noCheckTokenMap[c.Request.URL.Path] {
		log.Debug(c, "checkOpenId Do Not Need Check URI:%v RETURN", c.Request.URL.Path)
		return
	}
	// 如果是静态文件，也不需要检查OpenId
	if len(c.Request.URL.Path) > 7 && c.Request.URL.Path[:7] == "/static" {
		log.Debug(c, "checkOpenId Do Not Need Check URI:%v RETURN", c.Request.URL.Path)
		return
	}

	// 获取OpenId
	openId := c.Request.Header.Get("Openid")
	// 检查获取OpenId内容是否正常
	if len(openId) == 0 {
		log.Debug(c, "checkOpenId GET OpenId:%v", openId)
		c.AbortWithStatusJSON(http.StatusOK, tools.StatusUnauthorized)
		return
	}

	userP := models.MGetUserByOpenId(c, openId)
	// 检查获取UID内容是否正常
	if userP == nil {
		log.Debug(c, "checkOpenId GET token:%v User Not Found", openId)
		c.AbortWithStatusJSON(http.StatusOK, tools.StatusUnauthorized)
		return
	}

	// 正确的UID内容，设置到GIN Context中
	c.Set(tools.UserKey, userP)
	log.Debug(c, "User:%v", gotools.Data2Str(userP))
}

var EncryptSign = "sign"
var EncryptTime = "timestamp"

// 将参数都获取出来，之后按照参数的字典序排
// 如  map{ "api":"sdfdfEDD", "key":"sadfasd", "goods":"34", "sign":"234adf", "timestamp":"3452345"}
// 字典序 api   goods   key timestamp 不包含sign字段
// 之后将对应的内容串联到一起 map[key1]+map[key2]+...+map[keyN] 如  sdfdfEDD34sadfasd3452345
// 将secret合并到最后  sdfdfEDD34sadfasd3452345 + {secret}
// 检查对应的md5是否和sign一样，返回结果  return md5(map[key1]+map[key2]+...+map[keyN]+secret) == sign
func _checkEncrypt(c *gin.Context) bool {
	// 获取application/json里面的数据
	params := make(map[string]string)
	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(&params); err != nil {
		log.Debug(c, "_checkEncrypt decode params ERROR:%v", err)
		return false
	}

	log.Debug(c, "_checkEncrypt body:%v", gotools.Data2Str(params))
	c.Set("BodyData", gotools.Data2Str(params)) // 这边设置，方便后面输出内容
	timeStr := params[EncryptTime]
	timestamp, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil || math.Abs(float64(timestamp-time.Now().Unix())) > conf.Config.FuncTimeSecs {
		log.Debug(c, "_checkEncrypt timestamp:%v", timeStr)
		return false
	}

	// 将对应的KEY进行排序
	keys := make([]string, 0)
	for key := range params {
		if key == EncryptSign {
			continue
		}
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 按照排序的内容，组织
	var mapStr string
	for _, key := range keys {
		mapStr += params[key]
	}
	// 后面加上secret的内容
	mapStr += conf.Config.EncryptSecret

	// 检查是否对应，返回结果
	newMd5 := gotools.Md5(mapStr)
	log.Debug(c, "_checkEncrypt newMd5:%v sign:%v", newMd5, params[EncryptSign])
	flag := params[EncryptSign] == newMd5

	if flag {
		// 将body内的数据读出来在写进去
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(gotools.Data2Str(params))))
	}

	return flag
}

func checkEncrypt(c *gin.Context) {
	if conf.Config.Service.Mode != gin.DebugMode && c.Request.Method == http.MethodPost && !_checkEncrypt(c) {
		log.Debug(c, "checkEncrypt Failed")
		c.AbortWithStatusJSON(http.StatusOK, tools.StatusUnauthorized)
		return
	}
}

// 输出Header
func outputHeader(c *gin.Context) {
	log.Debug(c, "Header:%v", gotools.Data2Str(c.Request.Header))
}
