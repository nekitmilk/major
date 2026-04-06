package rules

import (
	"major/internal/models"
	"major/internal/service"

	"github.com/sirupsen/logrus"
)

type RulesHandler struct {
	log     *logrus.Logger
	conf    *models.Config
	service *service.Service
}

func NewRulesHandler(log *logrus.Logger, conf *models.Config, service *service.Service) *RulesHandler {
	return &RulesHandler{
		log:     log,
		conf:    conf,
		service: service,
	}
}
