package health

import "github.com/sirupsen/logrus"

type HealthHandler struct {
	log *logrus.Logger
}

func NewHealthHandler(log *logrus.Logger) *HealthHandler {
	return &HealthHandler{
		log: log,
	}
}
