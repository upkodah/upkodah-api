package api

import (
	"github.com/gin-gonic/gin"
	"github.com/upkodah/upkodah-api/pkg/db"
	"github.com/upkodah/upkodah-api/pkg/env"
	"github.com/upkodah/upkodah-api/pkg/facility"
	"github.com/upkodah/upkodah-api/pkg/room"
	"log"
	"net/http"
	"os"
)

func RunAPI() {
	db.InitDB()
	autoMigrate()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/setting", getSetting)
		v1.GET("/rooms", getRooms)
		v1.GET("/room/info/:id", getInfo)
	}

	s := &http.Server{
		Addr:    ":" + os.Getenv(env.HTTPPort),
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
		&room.Info{},
	)
	alterQuery := "ALTER DATABASE " + os.Getenv(env.DBName) + " CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci"
	db.Conn.Exec(alterQuery)
	db.Conn.Exec("ALTER TABLE infos CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;")
}
