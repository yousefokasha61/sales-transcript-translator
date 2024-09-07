package log

import "github.com/sirupsen/logrus"

type LoggerConf interface {
	LogLevel() string
}

func NewLogger(conf LoggerConf) *logrus.Logger {
	logger := logrus.New()
	level := conf.LogLevel()
	switch level {
	case "PanicLevel":
		logger.SetLevel(logrus.PanicLevel)
	case "FatalLevel":
		logger.SetLevel(logrus.FatalLevel)
	case "ErrorLevel":
		logger.SetLevel(logrus.ErrorLevel)
	case "WarnLevel":
		logger.SetLevel(logrus.WarnLevel)
	case "InfoLevel":
		logger.SetLevel(logrus.InfoLevel)
	case "TraceLevel":
		logger.SetLevel(logrus.TraceLevel)
	default:
		logger.SetLevel(logrus.DebugLevel)
	}
	return logger
}
