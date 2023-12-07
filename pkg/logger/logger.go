package logger

import (
	"github.com/Campus-Hub/server/pkg/consts"
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

//func Setup() {
//	// 设置日志格式为json格式
//	log.SetFormatter(&log.JSONFormatter{})
//
//	// 设置将日志输出到指定文件（默认的输出为stderr,标准错误）
//	// 日志消息输出可以是任意的io.writer类型
//	logFile := ...
//	file, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
//	log.SetOutput(file)
//
//	// 设置只记录日志级别为warn及其以上的日志
//	log.SetLevel(log.WarnLevel)
//}

// Logger global logrus entity for logging.
var Logger = log.New()

func Setup() {
	// Write-only for log level higher that Warn.
	// for full level written: log.SetLevel(log.TraceLevel)
	log.SetLevel(log.WarnLevel)

	// Set the log format as json.
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "01-02 15:04:05",
	})

	// Set the output file
	// the log output can be any type implements io.writer.
	logFile := consts.LogFilePath
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		Logger.Fatalln("Log file Found Error: ", err)
	}

	//log.SetOutput(file)
	log.SetOutput(io.MultiWriter(colorable.NewColorableStdout(), colorable.NewNonColorable(file)))
}
