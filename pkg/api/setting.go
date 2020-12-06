package api

import (
	"github.com/gin-gonic/gin"
	"github.com/upkodah/upkodah-api/pkg/facility"
	"net/http"
)

func getSetting(c *gin.Context) {
	fs := make([]facility.Facility, 0, 10)
	if err := facility.SelectAll(&fs); err != nil {
		resErr500(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   fs,
	})
}
