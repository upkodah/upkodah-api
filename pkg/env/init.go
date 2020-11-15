package env

import (
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/config/")

	InitDefault()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("There is no config file in $HOME/config/")
	}

}

func InitDefault() {

	// Config DB
	viper.SetDefault(DBHostname, "localhost")
	viper.SetDefault(DBUser, "root")
	viper.SetDefault(DBPassword, "mypassword")
	viper.SetDefault(DBPort, "3306")
	viper.SetDefault(DBName, "upkodah")

	// Config API
	viper.SetDefault(HTTPPort, "80")
	viper.SetDefault(HTTPSPort, "443")
}
