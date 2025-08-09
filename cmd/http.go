package cmd

import (
	"Ewallet-grpc/helpers"
	"Ewallet-grpc/internal/api"
	"Ewallet-grpc/internal/services"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	healthCheckSvc := services.HealthCheck{}
	healthCheckApi := api.HealthCheck{
		HealthCheckServices: &healthCheckSvc,
	}

	r := gin.Default()
	r.GET("/ping", healthCheckApi.HealthCheckHandlerHttp)
	port := helpers.GetEnv("APP_PORT", "")
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
