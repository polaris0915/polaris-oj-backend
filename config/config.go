package config

var (
	App   *appConfig
	Mysql *mysqlConfig
)

type appConfig struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
}

type mysqlConfig struct {
	MYSQL_DB_NAME     string `mapstructure:"MYSQL_DB_NAME"`
	MYSQL_DB_PASSWORD string `mapstructure:"MYSQL_DB_PASSWORD"`
}
