package api

import (
	"github.com/gin-gonic/gin"
	"github.com/upkodah/upkodah-api/pkg/facility"
	"log"
	"net/http"
)

func getSetting(c *gin.Context) {
	f := []facility.Facility{}

	if err := facility.Init(&f); err != nil {
		log.Fatalf("Error in getSetting : %s\n", err)
	}

	println(gin.H{"f": f})

	c.JSON(http.StatusOK, gin.H{
		"facilities": f,
	})
}
