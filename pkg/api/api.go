package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/upkodah/upkodah-api/pkg/db"
	"github.com/upkodah/upkodah-api/pkg/env"
	"github.com/upkodah/upkodah-api/pkg/facility"
	"github.com/upkodah/upkodah-api/pkg/region"
	"github.com/upkodah/upkodah-api/pkg/room"
	"io"
	"log"
	"net/http"
	"os"
)

func RunAPI() {
	db.InitDB()
	autoMigrate()

	r := gin.Default()

	gin.DisableConsoleColor()

	f, err := os.Create("log.txt")

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal("Fail to Close Logging file")
		}
	}()

	if err != nil {
		log.Fatal("Fail to Open Logging file")
	}
	gin.DefaultWriter = io.MultiWriter(f)

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/setting", getSetting)
		v1.GET("/rooms", getRooms)
		v1.GET("/room/info/:id", getRoomInfo)

	}

	s := &http.Server{
		Addr:    ":" + viper.GetString(env.HTTPPort),
		Handler: r,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func autoMigrate() {
	db.Conn.AutoMigrate(
		&facility.Facility{},
		&room.Room{},
		&region.Goo{},
		&region.Dong{},
		&region.Grid{},
		&region.Search{},
	)
}
