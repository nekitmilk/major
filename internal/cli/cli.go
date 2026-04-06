package cli

import (
	"major/internal/cli/check"
	"major/internal/service"
)

type CheckHandler interface {
	CliHandler()
}

type Cli struct {
	CheckHandler
}

func NewCli(service *service.Service) *Cli {
	return &Cli{
		CheckHandler: check.NewCheckHandler(service),
	}

}
