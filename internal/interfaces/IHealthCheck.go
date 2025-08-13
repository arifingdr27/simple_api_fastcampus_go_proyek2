package interfaces

import "github.com/gin-gonic/gin"

type (
	IHealthCheckHandler interface {
		HealthCheckHandlerHttp(c *gin.Context)
	}
	IHealthCheckServices interface {
		HealthCheckServices() (string, error)
	}
)

type IHealthCheckRepo interface {
	// tes()
}
