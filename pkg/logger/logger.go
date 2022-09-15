package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewDevelopmentConfig()

	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Debug(message string, tags ...string) {
	log.Debug(message)
}

func Info(message string, tags ...string) {
	log.Info(message)
}

func Warn(message string, tags ...string) {
	log.Warn(message)
}

func Error(message string, err error, tags ...string) {
	log.Error(message, zap.Error(err))
}

func Panic(message string, err error, tags ...string) {
	log.Panic(message)
	panic(err)
}
