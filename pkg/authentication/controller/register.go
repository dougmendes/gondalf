package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implement yet",
	})

}
func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong!!",
	})

}
