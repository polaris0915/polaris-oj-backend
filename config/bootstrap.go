package config

import (
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

const CONF_DIR = "./config/"
const DEV_ENV = "dev"
const PROD_ENV = "prod"

func init() {
	vp := viper.New()
	_, file, _, _ := runtime.Caller(0)
	configPath, _ := filepath.Abs(filepath.Dir(file) + "/application." + DEV_ENV + ".yaml")
	vp.SetConfigFile(configPath)
	if err := vp.ReadInConfig(); err != nil {
		panic(err)
	}
	vp.UnmarshalKey("app", &App)
	vp.UnmarshalKey("mysql", &Mysql)
	vp.UnmarshalKey("jwt", &Jwt)
	vp.UnmarshalKey("session", &Session)
	vp.UnmarshalKey("log", &Log)

	// fmt.Printf("%+v\n", App)
	// fmt.Printf("%+v\n", Mysql)
	// fmt.Printf("%+v\n", Jwt)
	// fmt.Printf("%+v\n", Session)

}
