package error

import (
	"errors"
	"fmt"
)

var (
	InternalServiceError = errors.New("something went wrong")
	DomainError          = errors.New("domain error")

	CountryNotFound   = fmt.Errorf("country not found: %w", DomainError)
	UserAlreadyExists = fmt.Errorf("user already exists: %w", DomainError)
	UserNotFound      = fmt.Errorf("user not found: %w", DomainError)
)
