package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/api"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/repository"
	"ewallet-ums/internal/services"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	depedency := depedencyInject()

	r := gin.Default()

	r.GET("/ping", depedency.HealthCheckApi.HealthCheckHandlerHttp)

	userv1 := r.Group("/user/v1")
	userv1.POST("/register", depedency.RegisterApi.RegisterHandler)
	userv1.POST("/login", depedency.LoginApi.LoginHandlerHttp)
	userv1.DELETE("/logout", depedency.LogoutApi.LogoutHandler)

	port := helpers.GetEnv("APP_PORT", "")
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}

type Depedency struct {
	UserRepository interfaces.IUserRepository

	HealthCheckApi interfaces.IHealthCheckHandler
	RegisterApi    interfaces.IRegisterHandler
	LoginApi       interfaces.ILoginHandler
	LogoutApi      interfaces.ILogoutHandler
}

func depedencyInject() *Depedency {
	healthCheckService := &services.HealthCheck{}
	healthCheckAPI := api.HealthCheck{
		HealthCheckServices: healthCheckService,
	}

	repository := &repository.UserRepository{
		DB: helpers.DB,
	}
	registerSvc := services.RegisterService{
		UserRepo: repository,
	}
	registerAPI := api.Register{
		RegisterService: &registerSvc,
	}

	loginSvc := services.LoginService{
		UserRepo: repository,
	}

	LoginApi := api.LoginHandler{
		LoginService: &loginSvc,
	}

	LogoutService := services.LogoutService{
		UserRepo: repository,
	}

	LogoutApi := api.LogoutHandler{
		LogoutService: &LogoutService,
	}

	return &Depedency{
		RegisterApi:    &registerAPI,
		LoginApi:       &LoginApi,
		HealthCheckApi: &healthCheckAPI,
		LogoutApi:      &LogoutApi,
	}
}
