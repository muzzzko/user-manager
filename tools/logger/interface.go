package logger

import "github.com/go-openapi/strfmt"

type ILogger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)

	WithError(err error) ILogger
	WithPanic(err interface{}) ILogger
	WithNickname(nickname string) ILogger
	WithUserID(userID strfmt.UUID) ILogger
}
