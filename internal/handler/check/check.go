package check

import (
	"major/internal/models"
	"major/internal/service"

	"github.com/sirupsen/logrus"
)

type CheckHandler struct {
	log     *logrus.Logger
	conf    *models.Config
	service *service.Service
}

func NewCheckHandler(log *logrus.Logger, conf *models.Config, service *service.Service) *CheckHandler {
	return &CheckHandler{
		log:     log,
		conf:    conf,
		service: service,
	}
}
