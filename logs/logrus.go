package logs

import (
	"fmt"
	"github.com/jangozw/gin-api-common/configs"
	"github.com/jangozw/gin-api-common/consts"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var logger *logrus.Logger

func init() {
	if l, err := initLogrus(); err != nil {
		fmt.Println("init logrus failed " + err.Error())
	} else {
		logger = l
	}
}
func Logger() *logrus.Logger {
	return logger
}

func initLogrus() (*logrus.Logger, error) {
	logDir := configs.GetLogDir()
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}

	filePrefix := logDir + "/api"
	//view latest log info via api.log, history info in api.xxx.log
	latestLogFile := filePrefix + ".log"

	logClient := logrus.New()
	logClient.Out = src
	//logClient.Out = os.Stdout //stdout will output in console
	logClient.SetLevel(logrus.DebugLevel)

	logWriter, err := rotatelogs.New(
		filePrefix+".%Y%m%d%H.log",
		rotatelogs.WithLinkName(latestLogFile),   // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(1*24*time.Hour),    // 文件最大保存时间
		rotatelogs.WithRotationTime(1*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		return nil, err
	}
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
	}
	/*
		formatter := &logrus.JSONFormatter{
			//设置日志格式
			TimestampFormat: consts.TimeLayoutYmdHis,
			//PrettyPrint: true,

		}*/
	formatter := &logrus.TextFormatter{
		TimestampFormat: consts.TimeLayoutYmdHis,
	}
	lfHook := lfshook.NewHook(writeMap, formatter)
	logClient.AddHook(lfHook)
	return logClient, nil
}
