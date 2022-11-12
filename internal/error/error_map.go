package error

import (
	"context"
	"errors"
	"fmt"

	"github/user-manager/tools/logger"
)

const (
	internalErrorGroup = iota
	userErrorGroup
)

const (
	errorCodeMask = "%03d-%03d"
)

func init() {
	checkMap := make(map[string]struct{})
	for _, code := range errorToCodeMap {
		if _, ok := checkMap[code]; ok {
			panic("errorToCodeMap contains duplicates")
		}

		checkMap[code] = struct{}{}
	}
}

func GetDomainErrCode(ctx context.Context, err error) string {
	var unwrapedErr = err
	for unwrapedErr != nil {
		if code, ok := errorToCodeMap[unwrapedErr]; ok {
			return code
		}

		unwrapedErr = errors.Unwrap(unwrapedErr)
	}

	logger.GetFromContext(ctx).WithError(err).Error("known error")

	return fmt.Sprintf(errorCodeMask, internalErrorGroup, unknownErrorCode)
}

func GetInternalServiceErrCode() string {
	return fmt.Sprintf(errorCodeMask, internalErrorGroup, internalServiceErrorCode)
}

func GetMethodNotAllowedErrCode() string {
	return fmt.Sprintf(errorCodeMask, internalErrorGroup, methodNotAllowedErrorCode)
}

func GetValidationErrCode() string {
	return fmt.Sprintf(errorCodeMask, internalErrorGroup, ValidationErrorCode)
}

var errorToCodeMap = map[error]string{
	UserAlreadyExists: fmt.Sprintf(errorCodeMask, userErrorGroup, userAlreadyExistsErrorCode),
	CountryNotFound:   fmt.Sprintf(errorCodeMask, userErrorGroup, countryNotFoundErrorCode),
	UserNotFound:      fmt.Sprintf(errorCodeMask, userErrorGroup, userNotFoundErrorCode),
}
