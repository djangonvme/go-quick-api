package logger

import (
	"fmt"
	"gitlab.com/qubic-pool/pkg/util"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Instance *logrus.Logger

/*type Logger struct {
	*logrus.Logger
}*/

func InitLogger(dir string) error {
	if Instance != nil {
		return nil
	}
	a, err := newLogger(dir)
	if err != nil {
		return nil
	}
	Instance = a
	return nil
}
func newLogger(dir string) (*logrus.Logger, error) {
	if dir == "" {
		dir = "/dev/null"
	}
	// view latest log info via api.log, history info in api.xxx.log
	logFile := fmt.Sprintf("%s/%s.log", dir, util.DateInt())
	client := logrus.New()
	// client.Out = src
	client.Out = os.Stdout
	// client.Out = os.Stdout //stdout will output in console
	client.SetLevel(logrus.DebugLevel)
	writer, err := rotatelogs.New(
		dir+"/.%Y%m%d%H.log",
		rotatelogs.WithLinkName(logFile),          // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(180*24*time.Hour),   // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		return nil, err
	}

	// formatter := &logrus.TextFormatter{
	//	TimestampFormat: consts.TimeLayoutYmdHis,
	//}

	client.AddHook(
		lfshook.NewHook(
			lfshook.WriterMap{
				logrus.InfoLevel:  writer,
				logrus.WarnLevel:  writer,
				logrus.ErrorLevel: writer,
				logrus.DebugLevel: writer,
				logrus.FatalLevel: writer,
				logrus.PanicLevel: writer,
				logrus.TraceLevel: writer,
			},
			&logrus.JSONFormatter{
				// 设置日志格式
				TimestampFormat: "2006-01-02 15:04:05",
				//PrettyPrint:     true,
			}),
	)

	client.Infof("new logger success, log dir: %v", dir)

	return client, nil
}
