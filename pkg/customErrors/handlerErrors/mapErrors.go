package handlerErrors

import (
	"errors"
	"major/pkg/customErrors/serviceErrors/serviceErrorsModels"
	"net/http"
)

func MapError(err error) int {
	if err != nil {
		switch {
		case errors.Is(err, serviceErrorsModels.ErrNotFound):
			return http.StatusNotFound
		case errors.Is(err, serviceErrorsModels.ErrDuplicate):
			return http.StatusConflict
		case errors.Is(err, serviceErrorsModels.ErrForeignKey):
			return http.StatusConflict
		case errors.Is(err, serviceErrorsModels.ErrGenerateUuid):
			return http.StatusConflict
		case errors.Is(err, serviceErrorsModels.ErrMultipleFound):
			return http.StatusConflict
		case errors.Is(err, serviceErrorsModels.ErrParsing):
			return http.StatusBadRequest
		case errors.Is(err, serviceErrorsModels.ErrConnection):
			return http.StatusInternalServerError
		case errors.Is(err, serviceErrorsModels.ErrUnknown):
			return http.StatusConflict
		case errors.Is(err, serviceErrorsModels.ErrCelExpression):
			return http.StatusBadRequest
		case errors.Is(err, serviceErrorsModels.ErrValidation):
			return http.StatusBadRequest
		default:
			return http.StatusInternalServerError
		}
	}

	return 0
}
