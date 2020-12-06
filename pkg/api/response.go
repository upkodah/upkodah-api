package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func resErr404(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  "error",
		"message": msg,
	})
}

func resOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}

func resErr404InvalidParam(c *gin.Context) {
	resErr404(c, "invalid param")
}

func resErr500(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  "error",
		"message": "internal error occur",
	})
}
