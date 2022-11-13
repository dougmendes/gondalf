package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewProject(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "new project",
	})

}
