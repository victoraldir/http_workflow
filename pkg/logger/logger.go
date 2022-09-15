package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger
var config zap.Config

func init() {
	var err error

	config := zap.NewDevelopmentConfig()

	// enconderConfig := zap.NewProductionEncoderConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// enconderConfig.TimeKey = "time"
	//enconderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//enconderConfig.StacktraceKey = ""

	//config.EncoderConfig = enconderConfig
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func SetLogLevel(level string) {

	switch level {
	case "debug":
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		config.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		config.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "panic":
		config.Level = zap.NewAtomicLevelAt(zap.PanicLevel)
	case "fatal":
		config.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
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
