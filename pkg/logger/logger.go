package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var (
	logger *logrus.Logger
)

func init() {
	logger = setLogger()
}

func setLogger() *logrus.Logger {
	l := logrus.StandardLogger()
	l.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339Nano,
		DisableQuote:    true,
	})
	l.SetOutput(os.Stdout)
	return l
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func With(key string, value interface{}) *logrus.Entry {
	return logger.WithField(key, value)
}
