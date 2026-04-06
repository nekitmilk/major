package main

import (
	"fmt"
	"major/pkg/customErrors/externalServiceErrors/externalServiceErrorsModels"
	"major/pkg/customErrors/externalServiceErrors/postgres"
	"major/pkg/customErrors/handlerErrors"
	"major/pkg/customErrors/serviceErrors"
	"major/pkg/customErrors/serviceErrors/serviceErrorsModels"
)

func main() {
	rep := repo{}
	servic := service{
		repo: rep,
	}
	handle := handler{
		service: servic,
	}

	for i := 0; i < 7; i++ {
		println(handle.yaHandler(i))
	}
}

// =====================================================
//
//	Handler
//
// =====================================================
type handler struct {
	service service
}

func (h handler) yaHandler(input int) (int, string) {
	errService := h.service.yaServiceMethod(input)

	if errService != nil {
		message := fmt.Sprintf("Your custom error message is: %s", errService.Error())
		code := handlerErrors.MapError(errService)

		return code, message
	}

	return 200, "All good"
}

// =====================================================
//                        Service
// =====================================================

type service struct {
	repo repo
}

func (s *service) yaServiceMethod(input int) error {

	switch input {
	case 5:
		return serviceErrorsModels.NewParseStringAsUuidError(fmt.Errorf("uuid error"), "")
	}

	errRepo := s.repo.yaRepoMethod(input)
	errService := serviceErrors.MapError(errRepo)

	return errService
}

// =====================================================
//                        Repository
// =====================================================

type repo struct{}

func (r *repo) yaRepoMethod(input int) error {
	errDb := dbQuery(input)

	err := postgres.MapError(errDb)

	if input == 6 {
		return externalServiceErrorsModels.NewNotFoundError(fmt.Errorf("entity in postgres"), "")
	}
	//err := clickhouse.MapError(errDb)
	if err != nil {
		return err
	}

	return nil
}

func dbQuery(input int) error {
	switch input {
	case 1:
		return fmt.Errorf("connection refused")
	case 2:
		return fmt.Errorf("err from postgres 23505")
	case 3:
		return fmt.Errorf("err from postgres 23503")
	case 4:
		return fmt.Errorf("err from postgres unknown")
	default:
		return nil
	}
}
