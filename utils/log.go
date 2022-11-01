package utils

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

var Log *logrus.Logger

func InitialiseLogger() error {
	var err error
	Log, err = NewLogger()
	if err!=nil{
		fmt.Println(err)
		return err
	}
	return nil
}

func NewLogger() (*logrus.Logger, error) {
	logFilePath := "logs/app"
	logFileName := "log"

	fileName := path.Join(logFilePath, logFileName)

	log := logrus.New()

	log.SetReportCaller(true)


	//Set rotatelogs
	writeMap := createLogWriter(fileName)

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		CallerPrettyfier: caller(),
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyFile: "caller",
		},
		TimestampFormat:"2006-01-02 15:04:05",
	})

	//New hook
	log.AddHook(lfHook)



	return log, nil
}


func caller() func(*runtime.Frame) (function string, file string) {
	return func(f *runtime.Frame) (function string, file string) {
		p, _ := os.Getwd()

		return "", fmt.Sprintf("%s:%d", strings.TrimPrefix(f.File, p), f.Line)
	}
}
type Logwriter map[logrus.Level]io.Writer

func createLogWriter(fileName string )  lfshook.WriterMap {
	//Set rotatelogs
	logWriter, _ := rotatelogs.New(
		//Split file name
		fileName + ".%Y-%m-%d.log",

		//Generate soft chain, point to the latest log file

		//Set maximum save time (7 days)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		//Set log cutting interval (1 day)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	return writeMap
}