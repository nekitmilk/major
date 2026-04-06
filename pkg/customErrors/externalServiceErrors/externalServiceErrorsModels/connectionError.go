package externalServiceErrorsModels

import (
	"errors"
	"fmt"
)

var ErrConnection = errors.New("failed connect to db")

type ConnectionError struct {
	originalError error
	comment       string
}

func NewConnectionError(originalError error, comment string) error {
	if originalError == nil {
		return nil
	}
	return &ConnectionError{
		originalError: originalError,
		comment:       comment,
	}
}

func (e *ConnectionError) Error() string {
	if e.comment != "" {
		return fmt.Sprintf(
			"%s: %s (%s)",
			ErrConnection.Error(),
			e.originalError.Error(),
			e.comment,
		)
	}
	return fmt.Sprintf(
		"%s: %s",
		ErrConnection.Error(),
		e.originalError.Error(),
	)
}

func (e *ConnectionError) Is(target error) bool {
	if target == ErrConnection {
		return true
	}
	var connectionError *ConnectionError
	ok := errors.As(target, &connectionError)
	return ok
}

func (e *ConnectionError) Unwrap() error {
	return e.originalError
}
