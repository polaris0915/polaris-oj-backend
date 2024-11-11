package config

import (
	// "fmt"

	"github.com/spf13/viper"
)

const CONF_DIR = "./config/"
const DEV_ENV = "dev"
const PROD_ENV = "prod"

func init() {
	vp := viper.New()
	vp.SetConfigFile(CONF_DIR + "application." + DEV_ENV + ".yaml")
	if err := vp.ReadInConfig(); err != nil {
		panic(err)
	}
	vp.UnmarshalKey("app", &App)
	vp.UnmarshalKey("mysql", &Mysql)

	// fmt.Printf("%+v\n", App)
	// fmt.Printf("%+v\n", Mysql)
}
