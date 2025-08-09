package api

import (
	"net/http"

	"Ewallet-grpc/helpers"
	"Ewallet-grpc/internal/interfaces"

	"github.com/gin-gonic/gin"
)

type HealthCheck struct {
	HealthCheckServices interfaces.IHealthCheckServices
}

func (api *HealthCheck) HealthCheckHandlerHttp(c *gin.Context) {
	msg, err := api.HealthCheckServices.HealthCheckServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	helpers.SendResponseHttp(c, http.StatusOK, msg, nil)
}
