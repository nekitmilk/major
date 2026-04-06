package serviceErrorsModels

import (
	"errors"
	"fmt"
)

var ErrDuplicate = errors.New("unique constraint error")

type DuplicateError struct {
	originalError error
	comment       string
}

func NewDuplicateError(originalError error, comment string) error {
	if originalError == nil {
		return nil
	}
	return &DuplicateError{
		originalError: originalError,
		comment:       comment,
	}
}

func (e *DuplicateError) Error() string {
	if e.comment != "" {
		return fmt.Sprintf(
			"%s: %s (%s)",
			ErrDuplicate.Error(),
			e.originalError.Error(),
			e.comment,
		)
	}
	return fmt.Sprintf(
		"%s: %s",
		ErrDuplicate.Error(),
		e.originalError.Error())
}

func (e *DuplicateError) Is(target error) bool {
	if target == ErrDuplicate {
		return true
	}
	var duplicateError *DuplicateError
	ok := errors.As(target, &duplicateError)
	return ok
}

func (e *DuplicateError) Unwrap() error {
	return e.originalError
}
