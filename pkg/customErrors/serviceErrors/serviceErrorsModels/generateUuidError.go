package serviceErrorsModels

import (
	"errors"
	"fmt"
)

var ErrGenerateUuid = errors.New("failed to generate uuid")

type GenerateUuidError struct {
	originalError error
	comment       string
}

func NewGenerateUuidError(originalError error, comment string) error {
	if originalError == nil {
		return nil
	}
	return &GenerateUuidError{
		originalError: originalError,
		comment:       comment,
	}
}

func (e *GenerateUuidError) Error() string {
	if e.comment != "" {
		return fmt.Sprintf(
			"%s: %s (%s)",
			ErrGenerateUuid.Error(),
			e.originalError.Error(),
			e.comment,
		)
	}
	return fmt.Sprintf(
		"%s: %s",
		ErrGenerateUuid.Error(),
		e.originalError.Error())
}

func (e *GenerateUuidError) Is(target error) bool {
	if target == ErrGenerateUuid {
		return true
	}
	var generateUuidError *GenerateUuidError
	ok := errors.As(target, &generateUuidError)
	return ok
}

func (e *GenerateUuidError) Unwrap() error {
	return e.originalError
}
