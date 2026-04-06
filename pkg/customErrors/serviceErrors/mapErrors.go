package serviceErrors

import (
	"errors"
	"major/pkg/customErrors/externalServiceErrors/externalServiceErrorsModels"
	"major/pkg/customErrors/serviceErrors/serviceErrorsModels"
)

func MapError(err error) error {
	if err != nil {
		switch {
		case errors.Is(err, externalServiceErrorsModels.ErrNotFound):
			return serviceErrorsModels.NewNotFoundError(err, "")
		case errors.Is(err, externalServiceErrorsModels.ErrForeignKey):
			return serviceErrorsModels.NewForeignKeyError(err, "")
		case errors.Is(err, externalServiceErrorsModels.ErrDuplicate):
			return serviceErrorsModels.NewDuplicateError(err, "")
		case errors.Is(err, externalServiceErrorsModels.ErrConnection):
			return serviceErrorsModels.NewConnectionError(err, "")
		case errors.Is(err, externalServiceErrorsModels.ErrValidateUuid):
			return serviceErrorsModels.NewValidationError(err, "")
		default:
			return externalServiceErrorsModels.NewUnknownError(err, "")
		}
	}

	return nil
}
