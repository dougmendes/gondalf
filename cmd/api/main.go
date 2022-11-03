package main

import (
	authentication "github.com/dougmendes/gondalf/pkg/authentication/controller"
	instances "github.com/dougmendes/gondalf/pkg/instances/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", instances.Pong)
	r.POST("/register", authentication.Register)
	r.Run()
}
