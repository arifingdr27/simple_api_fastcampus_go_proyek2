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

	ewalletv1 := r.Group("wallet/v1")
	ewalletv1.POST("/", depedency.WalletAPI.Create)

	port := helpers.GetEnv("APP_PORT", "")
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}

type Depedency struct {
	HealthCheckApi interfaces.IHealthCheckHandler
	WalletAPI      interfaces.IWalletAPI
}

func depedencyInject() *Depedency {
	healthCheckService := &services.HealthCheck{}
	healthCheckAPI := api.HealthCheck{
		HealthCheckServices: healthCheckService,
	}
	walletRepo := &repository.WalletRepo{
		DB: helpers.DB,
	}

	walletSvc := &services.WalletService{
		WalletRepo: walletRepo,
	}
	walletAPI := &api.WalletAPI{
		WalletService: walletSvc,
	}

	return &Depedency{
		HealthCheckApi: &healthCheckAPI,
		WalletAPI:      walletAPI,
	}
}
