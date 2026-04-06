package serviceErrorsModels

import (
	"errors"
	"fmt"
)

var ErrForeignKey = errors.New("foreign key constraint error")

type ForeignKeyError struct {
	originalError error
	comment       string
}

func NewForeignKeyError(originalError error, comment string) error {
	if originalError == nil {
		return nil
	}
	return &ForeignKeyError{
		originalError: originalError,
		comment:       comment,
	}
}

func (e *ForeignKeyError) Error() string {
	if e.comment != "" {
		return fmt.Sprintf(
			"%s: %s (%s)",
			ErrForeignKey.Error(),
			e.originalError.Error(),
			e.comment,
		)
	}
	return fmt.Sprintf(
		"%s: %s",
		ErrForeignKey.Error(),
		e.originalError.Error())
}

func (e *ForeignKeyError) Is(target error) bool {
	if target == ErrForeignKey {
		return true
	}
	var foreignKeyError *ForeignKeyError
	ok := errors.As(target, &foreignKeyError)
	return ok
}

func (e *ForeignKeyError) Unwrap() error {
	return e.originalError
}
