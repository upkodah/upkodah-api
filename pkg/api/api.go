package api

import (
	"github.com/gin-gonic/gin"
	"github.com/upkodah/upkodah-api/pkg/db"
	"net/http"
)

func RunAPI() {
	db.InitDB()
	db.Conn.AutoMigrate()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

}
