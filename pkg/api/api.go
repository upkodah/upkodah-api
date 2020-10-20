package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/upkodah/upkodah-api/pkg/db"
	"github.com/upkodah/upkodah-api/pkg/env"
	"log"
	"net/http"
)

func RunAPI() {
	db.InitDB()
	db.AutoMig()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	s := &http.Server{
		Addr:    ":" + viper.GetString(env.HTTPPort),
		Handler: r,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
