package check

import "major/internal/service"

type CheckHandler struct {
	service *service.Service
}

func NewCheckHandler(service *service.Service) *CheckHandler {
	return &CheckHandler{
		service: service,
	}
}
