package singleton

import (
	"fmt"
	"github.com/go-quick-api/config"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger(module string) func(cfg *config.Config) (*Logger, error) {
	return func(cfg *config.Config) (*Logger, error) {
		return newLogger(cfg, module)
	}
}
func newLogger(cfg *config.Config, module string) (*Logger, error) {
	/*	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return nil, err
		}*/
	logDir := cfg.General.LogDir
	if logDir == "" {
		logDir = "/tmp"
	}
	filePrefix := logDir + "/" + module
	// view latest log info via api.log, history info in api.xxx.log
	latestLogFile := filePrefix + ".log"

	logClient := logrus.New()
	// logClient.Out = src
	logClient.Out = os.Stdout
	// logClient.Out = os.Stdout //stdout will output in console
	logClient.SetLevel(logrus.DebugLevel)

	logWriter, err := rotatelogs.New(
		filePrefix+".%Y%m%d%H.log",
		rotatelogs.WithLinkName(latestLogFile),    // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(180*24*time.Hour),   // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		return nil, err
	}
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
	}
	formatter := &logrus.JSONFormatter{
		// 设置日志格式
		TimestampFormat: "2006-01-02 15:04:05",
		// PrettyPrint: true,

	}
	// formatter := &logrus.TextFormatter{
	//	TimestampFormat: consts.TimeLayoutYmdHis,
	//}
	lfHook := lfshook.NewHook(writeMap, formatter)
	logClient.AddHook(lfHook)

	fmt.Printf("new logger success, log dir: %v\n", logDir)
	logClient.Infof("new logger success, log dir: %v", logDir)

	return &Logger{logClient}, nil
}

func (l *Logger) NewDbLogger() *DBLogger {
	return &DBLogger{baseLog: l.Logger}
}
