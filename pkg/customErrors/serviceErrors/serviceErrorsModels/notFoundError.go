package serviceErrorsModels

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found error")

type NotFoundError struct {
	originalError error
	comment       string
}

func NewNotFoundError(originalError error, comment string) error {
	return &NotFoundError{
		originalError: originalError,
		comment:       comment,
	}
}

func (e *NotFoundError) Error() string {
	if e.comment != "" {
		return fmt.Sprintf(
			"%s: %s (%s)",
			ErrNotFound.Error(),
			e.originalError.Error(),
			e.comment,
		)
	}
	return fmt.Sprintf(
		"%s: %s",
		ErrNotFound.Error(),
		e.originalError.Error())
}

func (e *NotFoundError) Is(target error) bool {
	if target == ErrNotFound {
		return true
	}
	var notFoundError *NotFoundError
	ok := errors.As(target, &notFoundError)
	return ok
}
