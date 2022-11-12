package logger

import (
	"context"
)

const (
	loggerContextKey = "logger"
)

func EnrichContext(ctx context.Context, logger ILogger) context.Context {
	return context.WithValue(ctx, loggerContextKey, logger)
}

func GetFromContext(ctx context.Context) ILogger {
	logger, ok := ctx.Value(loggerContextKey).(ILogger)
	if !ok {
		panic("invalid logger in context")
	}

	return logger
}
