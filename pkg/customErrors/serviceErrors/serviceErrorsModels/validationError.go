package serviceErrorsModels

import (
	"errors"
	"fmt"
)

var ErrValidation = errors.New("validation error")

type ValidationError struct {
	originalError error
	comment       string
}

func NewValidationError(originalError error, comment string) error {
	if originalError == nil {
		return nil
	}
	return &ValidationError{
		originalError: originalError,
		comment:       comment,
	}
}

func (e *ValidationError) Error() string {
	if e.comment != "" {
		return fmt.Sprintf(
			"%s: %s (%s)",
			ErrValidation.Error(),
			e.originalError.Error(),
			e.comment,
		)
	}
	return fmt.Sprintf(
		"%s: %s",
		ErrValidation.Error(),
		e.originalError.Error())
}

func (e *ValidationError) Is(target error) bool {
	if target == ErrValidation {
		return true
	}
	var foreignKeyError *ValidationError
	ok := errors.As(target, &foreignKeyError)
	return ok
}

func (e *ValidationError) Unwrap() error {
	return e.originalError
}
