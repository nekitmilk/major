package main

import (
	"major/internal/cli"
	"major/internal/config"
	"major/internal/handler"
	"major/internal/repository"
	"major/internal/repository/sql"
	"major/internal/service"
	"major/pkg/logger"
	"major/pkg/server"

	"github.com/sirupsen/logrus"
)

// @title major - статический анализатор конфигов
// @version 0.0.1
// @description Документация для API major - статический анализатор конфигов
// @basePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	conf, errLoadConfig := config.LoadConfig()
	if errLoadConfig != nil {
		logrus.Fatalf("error load config file: %s \n", errLoadConfig.Error())
		return
	}

	log := logger.SetLogger(conf.Debug)
	config.LoadEnv(conf)

	db, errInitDB := sql.NewSQL(conf).InitDB()
	if errInitDB != nil {
		log.Fatalf("Error connect db: %s", errInitDB.Error())
	}
	//else {
	//	log.Infof("Successful connection to PostgreSQL (%s:%s)", conf.Db.Host, conf.Db.Port)
	//}
	repo := repository.NewRepository(db)
	services := service.NewService(conf, repo)

	if conf.ConsoleMode {
		cliHandler := cli.NewCli(services)
		cliHandler.CliHandler()
		return
	}

	handlers := handler.NewHandler(log, conf, services)

	srv := new(server.Server)

	errStartServer := srv.Run(conf.ServerSettings.PortHttp, handlers.InitRoutes(conf), conf.ServerSettings.ReadTimeout, conf.ServerSettings.WriteTimeout)
	if errStartServer != nil {
		log.Fatalf("Failed to start server: %s", errStartServer.Error())
	}

}
