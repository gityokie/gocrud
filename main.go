package main

import (
	"github.com/gityokie/gocrud/internal/logger"
	"github.com/gityokie/gocrud/internal/routers"
	"github.com/gityokie/gocrud/internal/store/mysql"
	"github.com/gityokie/gocrud/internal/utils"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
)

var err error

func main() {

	mysql, err := mysql.New()
	if err != nil {
		logger.Log.Fatal("failed to initialize database", err)
	}

	r := routers.SetupRouter()
	// running
	r.Run()
}

func init() {
	viper.AutomaticEnv()

	viper.SetEnvPrefix(utils.AppName)

	BindEnvs()

	viper.SetDefault(utils.Port, "8080")
	viper.SetDefault(utils.DbHost, "127.0.0.1")

	CheckMustBeSetEnvs()
}

func BindEnvs() {
	viper.BindEnv(utils.Port, "PORT")
	viper.BindEnv(utils.DbUser)
	viper.BindEnv(utils.DbPwd)
	viper.BindEnv(utils.DbPort)
	viper.BindEnv(utils.DbName)
	viper.BindEnv(utils.DbHost)
	viper.BindEnv(utils.DbTimeZone)
}

func EnvMustBeSet(key string) {
	if !viper.IsSet(key) {
		logger.Log.WithField(key, false).Fatal("not set")
	}
}

func CheckMustBeSetEnvs() {
	EnvMustBeSet(utils.DbUser)
	EnvMustBeSet(utils.DbPwd)
	EnvMustBeSet(utils.DbPort)
	EnvMustBeSet(utils.DbName)
	EnvMustBeSet(utils.DbTimeZone)
}
