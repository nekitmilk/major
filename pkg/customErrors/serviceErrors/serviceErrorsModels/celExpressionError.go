package serviceErrorsModels

import (
	"errors"
	"fmt"
)

var ErrCelExpression = errors.New("cel expression error")

type CelExpressionError struct {
	originalError error
	comment       string
}

func NewCelExpressionError(originalError error, comment string) error {
	if originalError == nil {
		return nil
	}
	return &CelExpressionError{
		originalError: originalError,
		comment:       comment,
	}
}

func (e *CelExpressionError) Error() string {
	if e.comment != "" {
		return fmt.Sprintf(
			"%s: %s (%s)",
			ErrCelExpression.Error(),
			e.originalError.Error(),
			e.comment,
		)
	}
	return fmt.Sprintf(
		"%s: %s",
		ErrCelExpression.Error(),
		e.originalError.Error())
}

func (e *CelExpressionError) Is(target error) bool {
	if target == ErrCelExpression {
		return true
	}
	var foreignKeyError *CelExpressionError
	ok := errors.As(target, &foreignKeyError)
	return ok
}

func (e *CelExpressionError) Unwrap() error {
	return e.originalError
}
