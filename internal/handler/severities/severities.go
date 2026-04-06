package severities

import (
	"major/internal/models"
	"major/internal/service"

	"github.com/sirupsen/logrus"
)

type SeveritiesHandler struct {
	log     *logrus.Logger
	conf    *models.Config
	service *service.Service
}

func NewSeveritiesHandler(log *logrus.Logger, conf *models.Config, service *service.Service) *SeveritiesHandler {
	return &SeveritiesHandler{
		log:     log,
		conf:    conf,
		service: service,
	}
}
