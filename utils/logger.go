package utils

import (
	"context"
	"google.golang.org/appengine/log"
	"github.com/sirupsen/logrus"
)

type LogType int

const (
	AppEngineLogType LogType = iota // for AppEngineLogType, default
	Logrus
)

type CustomLogger struct {
	LogType
}

func NewCustomLogger(logType LogType) CustomLogger {
	return CustomLogger{logType}
}

func (l *CustomLogger) Info(ctx context.Context, format string, args ...interface{}) {
	switch l.LogType {
	case AppEngineLogType:
		log.Infof(ctx, format, args)
	case Logrus:
		logrus.Info(args)
	}
}

func (l *CustomLogger) Error(ctx context.Context, format string, args ...interface{}) {
	switch l.LogType {
	case AppEngineLogType:
		log.Errorf(ctx, format, args)
	case Logrus:
		logrus.Error(args)
	}
}
