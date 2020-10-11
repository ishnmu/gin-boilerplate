package routes

import (
	healthcontroller "gin-boilerplate/controller/health"

	"github.com/gin-gonic/gin"
)

// Initialize routes
func Initialize(api *gin.Engine) {

	health := healthcontroller.NewHealthController()
	api.GET("/health", health.Check)
}
