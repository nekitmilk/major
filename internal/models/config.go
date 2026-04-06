package models

type Config struct {
	Debug          bool `mapstructure:"debug"`
	ConsoleMode    bool `mapstructure:"console_mode"`
	ServerSettings struct {
		PortHttp     string `mapstructure:"port_http"`
		ReadTimeout  int    `mapstructure:"read_timeout"`
		WriteTimeout int    `mapstructure:"write_timeout"`
	} `mapstructure:"server_settings"`
	Db struct {
		DriverName string `mapstructure:"driver_name"`
		Host       string `mapstructure:"host"`
		Port       string `mapstructure:"port"`
		User       string `mapstructure:"user"`
		Password   string `mapstructure:"password"`
		DbName     string `mapstructure:"dbName"`
		Sslmode    string `mapstructure:"ssl_mode"`
	} `mapstructure:"db"`
	EnvName struct {
		Debug       string `mapstructure:"debug"`
		ConsoleMode string `mapstructure:"console_mode"`
		PortHttp    string `mapstructure:"port_http"`
		DriverName  string `mapstructure:"driverName"`
		Host        string `mapstructure:"host"`
		Port        string `mapstructure:"port"`
		User        string `mapstructure:"user"`
		Password    string `mapstructure:"password"`
		DbName      string `mapstructure:"dbName"`
		Sslmode     string `mapstructure:"sslMode"`
	} `mapstructure:"env_name"`
}
