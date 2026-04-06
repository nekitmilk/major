package issues

import (
	"major/internal/models"
	"major/internal/service"

	"github.com/sirupsen/logrus"
)

type IssuesHandler struct {
	log     *logrus.Logger
	conf    *models.Config
	service *service.Service
}

func NewIssuesHandler(log *logrus.Logger, conf *models.Config, service *service.Service) *IssuesHandler {
	return &IssuesHandler{
		log:     log,
		conf:    conf,
		service: service,
	}
}
