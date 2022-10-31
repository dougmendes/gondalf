package main

import (
	. "github.com/dougmendes/gondalf/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", Pong)
	r.Run()
}
