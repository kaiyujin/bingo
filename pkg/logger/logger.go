package logger

import (
	bingoConfig "bingo/config"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func init() {
	var err error
	var config zap.Config

	if bingoConfig.IsLocal() {
		config = NewLoggerDevelopmentConfig()
	} else {
		config = NewLoggerProductionConfig()
	}
	l, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	defer l.Sync()
	log = l.Sugar()
}

func Info(args ...interface{}) {
	log.Info(args)
}

func Debug(args ...interface{}) {
	log.Debug(args)
}

func Error(args ...interface{}) {
	log.Error(args)
}
