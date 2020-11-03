package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"github.com/upkodah/upkodah-api/pkg/env"
	"github.com/upkodah/upkodah-api/pkg/facility"
	"github.com/upkodah/upkodah-api/pkg/region"
	"github.com/upkodah/upkodah-api/pkg/room"
	"log"
)

var (
	Conn *gorm.DB
)

func InitDB() {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString(env.DBUser),
		viper.GetString(env.DBPassword),
		viper.GetString(env.DBHostname),
		viper.GetString(env.DBPort),
		viper.GetString(env.DBName),
	)
	var err error
	fmt.Printf("db url : %s", dbURI)
	Conn, err = gorm.Open("mysql", dbURI)
	if err != nil {
		log.Fatal("Got error when connect database")
	}

	Conn.DB().SetMaxIdleConns(0)

	Conn.Exec(fmt.Sprintf("ALTER DATABASE %s DEFAULT CHARACTER SET utf8", viper.GetString(env.DBName)))
}

func AutoMig() {
	Conn.AutoMigrate(
		&facility.Facility{},
		&room.Room{},
		&region.Goo{},
		&region.Dong{},
		&region.Grid{},
		&region.Search{},
	)
}
