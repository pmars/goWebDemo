package log

import (
	"fmt"

	"goWebDemo/tools"

	"github.com/gin-gonic/gin"
	"github.com/pmars/beego/logs"
)

func Error(c *gin.Context, format string, params ...interface{}) {
	requestId, _ := c.Get(tools.RequestKey)
	logs.Error("RequestId:%v %v", requestId, fmt.Sprintf(format, params...))
}

func Debug(c *gin.Context, format string, params ...interface{}) {
	requestId, _ := c.Get(tools.RequestKey)
	logs.Debug("RequestId:%v %v", requestId, fmt.Sprintf(format, params...))
}

func Info(c *gin.Context, format string, params ...interface{}) {
	requestId, _ := c.Get(tools.RequestKey)
	logs.Info("RequestId:%v %v", requestId, fmt.Sprintf(format, params...))
}
