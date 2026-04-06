package externalServiceErrorsModels

import (
	"errors"
	"fmt"
)

// Нашли несколько сущностей, хотя должны были найти только одну
var ErrMultipleFound = errors.New("multiple found")

type MultipleFoundError struct {
	originalError error
	comment       string
}

func NewMultipleFoundError(originalError error, comment string) error {
	return &MultipleFoundError{
		originalError: originalError,
		comment:       comment,
	}
}

func (e *MultipleFoundError) Error() string {
	if e.comment != "" {
		return fmt.Sprintf(
			"%s: %s (%s)",
			ErrMultipleFound.Error(),
			e.originalError.Error(),
			e.comment,
		)
	}
	return fmt.Sprintf(
		"%s: %s",
		ErrMultipleFound.Error(),
		e.originalError)
}

func (e *MultipleFoundError) Is(target error) bool {
	if target == ErrMultipleFound {
		return true
	}
	var multipleFoundError *MultipleFoundError
	ok := errors.As(target, &multipleFoundError)
	return ok
}
