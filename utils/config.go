package utils

import (
	"github.com/spf13/viper"
	"os"
)

var (
	AppName        string
	AppHost        string
	AppPort        string
	AppEnvironment string
	AppLogLevel    string

	AppAuthUsername string
	AppAuthPassword string
)

func InitConfig() {
	// check app environment on env
	env := os.Getenv("APP_ENV")

	// check for value
	if env == "" {
		env = "development"
	}

	if env == "development" {
		// check for config.json
		viper.SetConfigFile(`config.json`)

		// read file
		err := viper.ReadInConfig()
		if err != nil {
			Fatal(err, "InitConfig", "")
			panic(err)
		}

		// get variable for app
		AppName = viper.GetString("app.name")
		AppHost = viper.GetString("app.host")
		AppPort = viper.GetString("app.port")
		AppEnvironment = viper.GetString("app.environment")
		AppLogLevel = viper.GetString("app.log.level")
		AppAuthUsername = viper.GetString("app.auth.username")
		AppAuthPassword = viper.GetString("app.auth.password")

		// create return
		return
	}

	if env == "staging" || env == "production" {
		// get variable for app
		AppName = os.Getenv("APP_NAME")
		AppPort = os.Getenv("APP_PORT")
		AppEnvironment = os.Getenv("APP_ENV")
		AppLogLevel = os.Getenv("APP_LOG_LEVEL")
	}
}
