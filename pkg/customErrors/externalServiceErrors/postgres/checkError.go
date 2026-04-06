package postgres

import (
	"strings"
)

func IsDuplicateError(err error) bool {
	if err == nil {
		return false
	}

	errMsg := strings.ToLower(err.Error())

	if strings.Contains(errMsg, "23505") {
		return true
	}

	return strings.Contains(errMsg, "unique constraint") ||
		strings.Contains(errMsg, "duplicate key") ||
		strings.Contains(errMsg, "duplicate entry") ||
		strings.Contains(errMsg, "already exists")
}

func IsForeignKeyError(err error) bool {
	if err == nil {
		return false
	}

	errMsg := strings.ToLower(err.Error())

	if strings.Contains(errMsg, "23503") {
		return true
	}

	return strings.Contains(errMsg, "foreign key constraint") ||
		strings.Contains(errMsg, "foreign key violation") ||
		strings.Contains(errMsg, "referential integrity") ||
		strings.Contains(errMsg, "violates foreign key")
}

func IsConnectionError(err error) bool {
	if err == nil {
		return false
	}

	errMsg := strings.ToLower(err.Error())

	if strings.Contains(errMsg, "connection refused") {
		return true
	}

	return false
}

func IsValidateUuidError(err error) bool {
	if err == nil {
		return false
	}

	errMsg := strings.ToLower(err.Error())

	if strings.Contains(errMsg, "22p02") || strings.Contains(errMsg, "invalid input syntax for type uuid") {
		return true
	}

	return false
}
