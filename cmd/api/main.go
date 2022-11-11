package main

import (
	instances "github.com/dougmendes/gondalf/pkg/instances/controller"
	projectGenerator "github.com/dougmendes/gondalf/pkg/project-generator/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", instances.Pong)
	r.POST("/project-generator", projectGenerator.NewProject)
	r.Run()
}
