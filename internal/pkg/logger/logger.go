package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func setup() *zap.Logger {
	encoder := zapcore.NewConsoleEncoder(
		zap.NewDevelopmentEncoderConfig(),
	)

	logger := zap.New(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.InfoLevel),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.WarnLevel),
		zap.Development(),
	)

	zap.RedirectStdLog(logger)

	zap.ReplaceGlobals(logger)

	return logger
}

var logger = setup() //nolint:gochecknoglobals

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Err(err error, msg string, fields ...zapcore.Field) {
	logger.With(zap.Error(err)).Error(msg, fields...)
}

func Info(msg string, fields ...zapcore.Field) {
	logger.Info(msg, fields...)
}

func StdAt(level zapcore.Level) *log.Logger {
	l, err := zap.NewStdLogAt(logger.WithOptions(zap.AddCallerSkip(-1)), level)
	if err != nil {
		panic(err)
	}
	return l
}
