package healthcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthController checks the health of the microservices
type HealthController interface {
	Check(ctx *gin.Context)
}

type healthController struct {
}

// Check returns 200 if the service is up and running
func (hc healthController) Check(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "gin-boilerplate is up and running",
	})
	return
}

// NewHealthController initializer
func NewHealthController() HealthController {
	return healthController{}
}
