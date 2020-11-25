package api

import (
	"github.com/gin-gonic/gin"
	"github.com/upkodah/upkodah-api/pkg/room"
	"log"
	"net/http"
	"strconv"
)

func getRooms(c *gin.Context) {
	rooms := make([]room.Room, 0, 10)
	room.SelectRooms(&rooms)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   rooms,
	})
}

func getInfo(c *gin.Context) {
	info := room.Info{}
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "please check info id",
		})
	}

	roomID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "invalid info id",
		})
	}

	info.RoomID = uint(roomID)

	if err := info.First(); err != nil {
		log.Printf("Error in getInfo : %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "invalid info id",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   info,
	})
}
