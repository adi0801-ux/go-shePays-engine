package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"time"
)

//Log to file
func LoggerToFile() fiber.Handler {

	logFilePath := "logs/api"
	logFileName := "log"

	////Log file
	fileName := path.Join(logFilePath, logFileName)

	log := logrus.New()

	writeMap := createLogWriter(fileName)

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})

	//New hook
	log.AddHook(lfHook)

	return func(c *fiber.Ctx) error {
		//Start time
		startTime := time.Now()

		//Process request
		c.Next()

		//End time
		endTime := time.Now()

		//Execution time
		latencyTime := endTime.Sub(startTime).Microseconds()

		//Request method
		reqMethod := c.Method()

		//Request routing
		reqUri := c.Path()

		// status code
		statusCode := c.Response().StatusCode()

		// request IP
		clientIP := c.IP()

		//Log format
		log.WithFields(logrus.Fields{
			"status_code"  : statusCode,
			"latency_time(ms)" : latencyTime,
			"client_ip"    : clientIP,
			"req_method"   : reqMethod,
			"req_uri"      : reqUri,
		}).Info()
		return nil
	}
}

