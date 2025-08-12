package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/api"
	"ewallet-ums/internal/repository"
	"ewallet-ums/internal/services"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	depedency := depedencyInject()

	r := gin.Default()

	r.GET("/ping", depedency.HealthCheckAPI.HealthCheckHandlerHttp)

	userv1 := r.Group("/user/v1")
	userv1.POST("/register", depedency.RegisterAPI.RegisterHandler)
	userv1.POST("/login", depedency.LoginApi.LoginHandlerHttp)

	port := helpers.GetEnv("APP_PORT", "")
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}

type Depedency struct {
	HealthCheckAPI api.HealthCheck
	RegisterAPI    api.Regsiter
	LoginApi       api.LoginHandler
}

func depedencyInject() Depedency {
	healthCheckService := services.HealthCheck{}
	healthCheckAPI := api.HealthCheck{
		HealthCheckServices: &healthCheckService,
	}

	repository := &repository.UserRepository{
		DB: helpers.DB,
	}
	registerSvc := services.RegisterService{
		UserRepo: repository,
	}
	registerAPI := api.Regsiter{
		RegisterService: &registerSvc,
	}

	loginSvc := services.LoginService{
		UserRepo: repository,
	}

	LoginApi := api.LoginHandler{
		LoginService: &loginSvc,
	}

	return Depedency{
		RegisterAPI:    registerAPI,
		LoginApi:       LoginApi,
		HealthCheckAPI: healthCheckAPI,
	}
}
