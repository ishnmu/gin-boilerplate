package main

//go:generate mockgen -source controller/health/health_controller.go -destination mocks/controller/mock_health_controller.go -package mock_health_controller

import (
	"github.com/gin-gonic/gin"
	"gin-boilerplate/routes"
)

func main() {
	api := gin.Default()
	routes.Initialize(api)
	api.Run()
}
