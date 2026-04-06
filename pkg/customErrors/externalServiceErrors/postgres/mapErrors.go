package postgres

import (
	"major/pkg/customErrors/externalServiceErrors/externalServiceErrorsModels"
)

func MapError(err error) error {
	if err != nil {
		if IsDuplicateError(err) {
			return externalServiceErrorsModels.NewDuplicateError(err, "")
		} else if IsForeignKeyError(err) {
			return externalServiceErrorsModels.NewForeignKeyError(err, "")
		} else if IsConnectionError(err) {
			return externalServiceErrorsModels.NewConnectionError(err, "")
		} else if IsValidateUuidError(err) {
			return externalServiceErrorsModels.NewValidateUuidError(err, "")
		}

		return externalServiceErrorsModels.NewUnknownError(err, "")
	}

	return nil
}
