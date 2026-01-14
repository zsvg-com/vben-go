package logs

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"vben/app/admin/model/constants"
	"vben/app/admin/model/monitor"
	"vben/app/core/utils"
	"vben/config"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// get 方法写入log 文档里面
// post put del 方法 写入数据库
// 错误日志写入数据库
func Logger() gin.HandlerFunc {
	if !config.LogConfig.Enabled {
		return gin.Logger()
	}
	mode := config.LogConfig.LogMode
	switch mode {
	case "default":
		return gin.Logger()
	case "mysql":
		return MysqlLogger()
	case "file":
		return Logrus()
	case "es":
		return EsLogger()
	}
	return gin.Logger()
}

func Logrus() gin.HandlerFunc {
	filePath := config.LogConfig.FilePath

	scr, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		utils.ExistDir(filePath)
	}
	logger := logrus.New()

	logger.Out = scr

	logger.SetLevel(logrus.DebugLevel)

	logWriter, _ := retalog.New(
		filePath+"%Y%m%d.log",
		retalog.WithMaxAge(7*24*time.Hour),
		retalog.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: constants.TimeFormat,
	})

	logger.AddHook(Hook)

	return func(context *gin.Context) {
		startTime := time.Now()
		context.Next()
		// 结束时间
		endTime := time.Now()
		stopTime := time.Since(startTime).Milliseconds()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		println(latencyTime)
		spendTime := fmt.Sprintf("%d ms", stopTime)
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := context.Writer.Status()
		clientIp := context.ClientIP()
		userAgent := context.Request.UserAgent()
		dataSize := context.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := context.Request.Method
		path := context.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})
		if len(context.Errors) > 0 {
			entry.Error(context.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}

func MysqlLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		startTime := time.Now()
		context.Next()
		// 结束时间
		endTime := time.Now()
		stopTime := time.Since(startTime).Milliseconds()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		println(latencyTime)
		spendTime := fmt.Sprintf("%d ms", stopTime)
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := context.Writer.Status()
		clientIp := context.ClientIP()
		userAgent := context.Request.UserAgent()
		dataSize := context.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := context.Request.Method
		path := context.Request.RequestURI
		//	写入数据库
		println(spendTime)
		println(hostName)
		println(statusCode)
		println(clientIp)
		println(userAgent)
		println(method)
		println(path)
		if method != "GET" && method != "OPTIONS" {
			if showLog(path) {
				var operLog = monitor.SysOperLog{
					Title:        "操作日志",
					BusinessType: monitor.BusinessTypeOther,
					Method:       "" + method,
					OperatorType: monitor.OperatorTypeAdmin,
					OperParam:    "",
					JsonResult:   "",
					Status:       "" + strconv.Itoa(statusCode),
					OperUrl:      path,
					OperIp:       clientIp,
				}
				operLog.OperationLogAdd(context)
			}
		}
	}
}

func EsLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		startTime := time.Now()
		context.Next()
		// 结束时间
		endTime := time.Now()
		stopTime := time.Since(startTime).Milliseconds()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		println(latencyTime)
		spendTime := fmt.Sprintf("%d ms", stopTime)
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := context.Writer.Status()
		clientIp := context.ClientIP()
		userAgent := context.Request.UserAgent()
		dataSize := context.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := context.Request.Method
		path := context.Request.RequestURI
		//	写入es
		println(spendTime)
		println(hostName)
		println(statusCode)
		println(clientIp)
		println(userAgent)
		println(method)
		println(path)
	}
}

func showLog(path string) bool {
	earte := config.LogConfig.Filtered
	for _, s := range earte {
		if s == path {
			return false
		}
	}
	return true
}

// poster logo
func Poster() {
	fg := color.New(color.FgBlue)
	logo := `
  _____                            _             _____         
 |  __ \                          (_)           / ____|        
 | |__) |  _   _    ___    _   _   _   ______  | |  __    ___  
 |  _  /  | | | |  / _ \  | | | | | | |______| | | |_ |  / _ \ 
 | | \ \  | |_| | | (_) | | |_| | | |          | |__| | | (_) |
 |_|  \_\  \__,_|  \___/   \__, | |_|           \_____|  \___/ 
                            __/ |                              
                           |___/                               
` +
		"Author:		vben\r\n" +
		"Version:	" + config.ProjectVersion + "\r\n" +
		"MiniGO_SDK: 	" + config.MinGoVersion + "\r\n" +
		"Link: https://gitee.com/vben/vben-go"
	fg.Println(logo)
}
