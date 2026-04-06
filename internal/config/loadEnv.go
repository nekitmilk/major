package config

import (
	"major/internal/models"
	"os"
	"strconv"
)

func LoadEnv(conf *models.Config) {
	debug := os.Getenv(conf.EnvName.Debug)
	if debug != "" {
		conf.Debug, _ = strconv.ParseBool(debug)
	}

	consoleMode := os.Getenv(conf.EnvName.ConsoleMode)
	if consoleMode != "" {
		conf.ConsoleMode, _ = strconv.ParseBool(consoleMode)
	}

	portHttp := os.Getenv(conf.EnvName.PortHttp)
	if portHttp != "" {
		conf.ServerSettings.PortHttp = portHttp
	}

	driverName := os.Getenv(conf.EnvName.DriverName)
	if driverName != "" {
		conf.Db.DriverName = driverName
	}

	dbHost := os.Getenv(conf.EnvName.Host)
	if dbHost != "" {
		conf.Db.Host = dbHost
	}

	dbPort := os.Getenv(conf.EnvName.Port)
	if dbPort != "" {
		conf.Db.Port = dbPort
	}

	dbUser := os.Getenv(conf.EnvName.User)
	if dbUser != "" {
		conf.Db.User = dbUser
	}

	dbPassword := os.Getenv(conf.EnvName.Password)
	if dbPassword != "" {
		conf.Db.Password = dbPassword
	}

	dbName := os.Getenv(conf.EnvName.DbName)
	if dbName != "" {
		conf.Db.DbName = dbName
	}

	dbSSLMode := os.Getenv(conf.EnvName.Sslmode)
	if dbSSLMode != "" {
		conf.Db.Sslmode = dbSSLMode
	}
}
