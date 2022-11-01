package main

import (
	. "github.com/dougmendes/gondalf/pkg/instances/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", Pong)
	r.Run()
}
