package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/upkodah/upkodah-api/pkg/env"
	"log"
	"os"
)

var (
	Conn *gorm.DB
)

func InitDB() {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv(env.DBUser),
		os.Getenv(env.DBPassword),
		os.Getenv(env.DBHost),
		os.Getenv(env.DBPort),
		os.Getenv(env.DBName),
	)
	log.Printf("db url : %s", dbURI)

	conn, err := gorm.Open("mysql", dbURI)
	if err != nil {
		log.Fatal("Got error when connect database")
	}

	Conn = conn

	Conn.DB().SetMaxIdleConns(0)

	Conn.Exec(fmt.Sprintf("ALTER DATABASE %s DEFAULT CHARACTER SET utf8", os.Getenv(env.DBName)))
}
