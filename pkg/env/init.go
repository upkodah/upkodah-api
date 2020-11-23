package env

import (
	"log"
	"os"
)

func InitConfig() {
	// Config DB
	setDefaultEnvStr(DBName, "upkodah")
	setDefaultEnvStr(DBHost, "localhost")
	setDefaultEnvStr(DBUser, "root")
	setDefaultEnvStr(DBPassword, "mypassword")
	setDefaultEnvStr(DBPort, "3306")

	// Config API
	setDefaultEnvStr(HTTPPort, "80")
	setDefaultEnvStr(HTTPSPort, "443")
}

func setDefaultEnvStr(key string, val string) {
	if os.Getenv(key) == "" {
		if err := os.Setenv(key, val); err != nil {
			log.Fatalf("Error in setDefaultEnvStr : %s\n", err)
		}
	}
}
