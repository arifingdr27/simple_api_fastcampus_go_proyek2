package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/api"
	"ewallet-ums/internal/repository"
	"ewallet-ums/internal/services"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	healthCheckSvc := services.HealthCheck{}
	healthCheckApi := api.HealthCheck{
		HealthCheckServices: &healthCheckSvc,
	}

	r := gin.Default()

	r.GET("/ping", healthCheckApi.HealthCheckHandlerHttp)

	registerRepository := &repository.RegisterRepository{
		DB: helpers.DB,
	}
	registerService := &services.RegisterService{
		RegisterRepo: registerRepository,
	}
	registerController := &api.Regsiter{
		RegisterService: registerService,
	}
	userv1 := r.Group("/user/v1")
	userv1.POST("/register", registerController.RegisterHandler)

	port := helpers.GetEnv("APP_PORT", "")
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
