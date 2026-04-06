package serviceErrorsModels

import (
	"errors"
	"fmt"
)

var ErrUnknown = errors.New("unknown error for service from db")

type UnknownError struct {
	originalError error
	comment       string
}

func NewUnknownError(originalError error, comment string) error {
	if originalError == nil {
		return nil
	}
	return &UnknownError{
		originalError: originalError,
		comment:       comment,
	}
}

func (e *UnknownError) Error() string {
	if e.comment != "" {
		return fmt.Sprintf(
			"%s: %s (%s)",
			ErrUnknown.Error(),
			e.originalError.Error(),
			e.comment,
		)
	}
	return fmt.Sprintf(
		"%s: %s",
		ErrUnknown.Error(),
		e.originalError.Error())
}

func (e *UnknownError) Is(target error) bool {
	if target == ErrUnknown {
		return true
	}
	var unknownError *UnknownError
	ok := errors.As(target, &unknownError)
	return ok
}

func (e *UnknownError) Unwrap() error {
	return e.originalError
}
