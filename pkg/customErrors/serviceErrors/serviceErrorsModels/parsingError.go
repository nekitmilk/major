package serviceErrorsModels

import (
	"errors"
	"fmt"
)

var ErrParsing = errors.New("parsing error")

type ParsingError struct {
	originalError error
	comment       string
}

func NewParsingError(originalError error, comment string) error {
	if originalError == nil {
		return nil
	}
	return &ParsingError{
		originalError: originalError,
		comment:       comment,
	}
}

func (e *ParsingError) Error() string {
	if e.comment != "" {
		return fmt.Sprintf(
			"%s: %s (%s)",
			ErrParsing.Error(),
			e.originalError.Error(),
			e.comment,
		)
	}
	return fmt.Sprintf(
		"%s: %s",
		ErrParsing.Error(),
		e.originalError.Error())
}

func (e *ParsingError) Is(target error) bool {
	if target == ErrParsing {
		return true
	}
	var parseStringAsUuidError *ParsingError
	ok := errors.As(target, &parseStringAsUuidError)
	return ok
}

func (e *ParsingError) Unwrap() error {
	return e.originalError
}
