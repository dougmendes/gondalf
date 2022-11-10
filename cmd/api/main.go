package main

import (
	instances "github.com/dougmendes/gondalf/pkg/instances/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", instances.Pong)
	r.Run()
}
