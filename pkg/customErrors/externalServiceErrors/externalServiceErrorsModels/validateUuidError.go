package externalServiceErrorsModels

import (
	"errors"
	"fmt"
)

var ErrValidateUuid = errors.New("invalid input uuid")

type ValidateUuidError struct {
	originalError error
	comment       string
}

func NewValidateUuidError(originalError error, comment string) error {
	if originalError == nil {
		return nil
	}
	return &ValidateUuidError{
		originalError: originalError,
		comment:       comment,
	}
}

func (e *ValidateUuidError) Error() string {
	if e.comment != "" {
		return fmt.Sprintf(
			"%s: %s (%s)",
			ErrValidateUuid.Error(),
			e.originalError.Error(),
			e.comment,
		)
	}

	return fmt.Sprintf(
		"%s: %s",
		ErrValidateUuid.Error(),
		e.originalError.Error())
}

func (e *ValidateUuidError) Is(target error) bool {
	if target == ErrValidateUuid {
		return true
	}
	var foreignKeyError *ValidateUuidError
	ok := errors.As(target, &foreignKeyError)
	return ok
}

func (e *ValidateUuidError) Unwrap() error {
	return e.originalError
}
