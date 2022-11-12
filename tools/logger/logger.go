package logger

import (
	"github.com/go-openapi/strfmt"
	"go.uber.org/zap"

	"github/user-manager/internal/constant"
)

type Logger struct {
	logger *zap.Logger
}

func CreateLogger(environment string) (*Logger, error) {
	var (
		zapl *zap.Logger
		err  error
	)

	if environment == constant.ProdEnv {
		zapl, err = zap.NewProduction()
	} else {
		devCfg := zap.NewDevelopmentConfig()
		devCfg.DisableStacktrace = true
		zapl, err = devCfg.Build()
	}
	if err != nil {
		return &Logger{}, err
	}

	return NewLogger(zapl), nil
}

func NewLogger(logger *zap.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l *Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *Logger) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l *Logger) Error(msg string) {
	l.logger.Error(msg)
}

func (l *Logger) Fatal(msg string) {
	l.logger.Fatal(msg)
}

func (l *Logger) WithError(err error) ILogger {
	zapl := l.logger.With(
		zap.Any("error", err),
	)

	return &Logger{logger: zapl}
}

func (l *Logger) WithNickname(nickname string) ILogger {
	zapl := l.logger.With(
		zap.String("nickname", nickname),
	)

	return &Logger{logger: zapl}
}

func (l *Logger) WithUserID(userID strfmt.UUID) ILogger {
	zapl := l.logger.With(
		zap.String("userID", userID.String()),
	)

	return &Logger{logger: zapl}
}

func (l *Logger) WithPanic(err interface{}) ILogger {
	zapl := l.logger.With(
		zap.Any("panic", err),
	)

	return &Logger{logger: zapl}
}
