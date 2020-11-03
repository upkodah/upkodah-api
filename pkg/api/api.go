package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/upkodah/upkodah-api/pkg/db"
	"github.com/upkodah/upkodah-api/pkg/env"
	"io"
	"log"
	"net/http"
	"os"
)

func RunAPI() {
	db.InitDB()
	db.AutoMig()

	r := gin.Default()

	gin.DisableConsoleColor()

	f, err := os.Create("log.txt")
	if err != nil {
		log.Fatal("Fail to Open Loggin App")
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
