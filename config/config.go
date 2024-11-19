package config

var (
	App     *appConfig
	Mysql   *mysqlConfig
	Jwt     *jwtConfig
	Session *sessionConfig
	Log     *logConfig
)

type appConfig struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
	Log  struct {
		FilePath         string `mapstructure:"path"`
		FileMaxSize      int    `mapstructure:"max_size"`
		BackupFileMaxAge int    `mapstructure:"max_age"`
	}
}

type mysqlConfig struct {
	MYSQL_DB_NAME     string `mapstructure:"MYSQL_DB_NAME"`
	MYSQL_DB_PASSWORD string `mapstructure:"MYSQL_DB_PASSWORD"`
}

type jwtConfig struct {
	ValidTime int64  `mapstructure:"validTime"`
	Key       string `mapstructure:"key"`
}

type sessionConfig struct {
	SESSION_PAIRKEY string `mapstructure:"SESSION_PAIRKEY"`
}

type ctxKey string
type logConfig struct {
	SESSION_PAIRKEY string `mapstructure:"SESSION_PAIRKEY"`
	LogPath         string `mapstructure:"LogPath"`
	LogContextKey   ctxKey `mapstructure:"LogContextKey"`
}
