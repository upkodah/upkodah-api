package api

import (
	"github.com/gin-gonic/gin"
	"github.com/upkodah/upkodah-api/pkg/bus"
	"github.com/upkodah/upkodah-api/pkg/db"
	"github.com/upkodah/upkodah-api/pkg/env"
	"github.com/upkodah/upkodah-api/pkg/facility"
	"github.com/upkodah/upkodah-api/pkg/metro"
	"github.com/upkodah/upkodah-api/pkg/room"
	"github.com/upkodah/upkodah-api/pkg/search"
	"log"
	"net/http"
	"os"
)

func RunAPI() {

	db.InitDB()
	autoMigrate()

	if err := metro.InitMetro(); err != nil {
		log.Printf("Error in RunAPI/%s\n", err)
	}
	if err := bus.InitBus(); err != nil {
		log.Printf("Error in RunAPI/%s\n", err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := r.Group("/v1")
	{
		// v1/setting
		v1.GET("/setting", getSetting)
		// v1/rooms
		v1.GET("/rooms", getRooms)
		// v1/rooms/:searchID
		v1.GET("/rooms/:search_id", getRoomsBySearchID)
		// v1/room/:id/info
		v1.GET("/room/:id/info", getInfo)
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
		&search.Search{},
		&search.Item{},
		&metro.Metro{},
	)

	// Setting DB Charset utf8mb4 for emoji
	alterQuery := "ALTER DATABASE " + os.Getenv(env.DBName) + " CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci"
	db.Conn.Exec(alterQuery)
	db.Conn.Exec("ALTER TABLE infos CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;")
}
