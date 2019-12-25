package logger

import (
	"bufio"
	"local/gin_init/common/config"
	"log"
	"os"
	"path"
	"time"

	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func init() {
	var logFormat logrus.Formatter
	if config.ServerConfig.Logger.LogFormat == "json" {
		logFormat = &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		}
	} else {
		logFormat = &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		}
	}
	logrus.SetFormatter(logFormat)

	baseLogPath := path.Join(config.ServerConfig.Logger.LogDir,
		config.ServerConfig.Logger.LogFileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		log.Println("config local file system logger errors")
	}

	switch config.ServerConfig.Logger.LogLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetOutput(os.Stderr)
	case "info":
		setNull()
		logrus.SetLevel(logrus.InfoLevel)
		logrus.SetOutput(os.Stderr)
	case "warn":
		setNull()
		logrus.SetLevel(logrus.WarnLevel)
	case "errors":
		setNull()
		logrus.SetLevel(logrus.ErrorLevel)
		// 显示行号等信息
		logrus.SetReportCaller(true)
	default:
		setNull()
		logrus.SetLevel(logrus.InfoLevel)
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, logFormat)
	logrus.AddHook(lfHook)
}

func setNull() {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println("err", err)
	}
	writer := bufio.NewWriter(src)
	log.SetOutput(writer)
}
