package cmd

import (
	"ewallet-wallet/external"
	"ewallet-wallet/helpers"
	"ewallet-wallet/internal/api"
	"ewallet-wallet/internal/interfaces"
	"ewallet-wallet/internal/repository"
	"ewallet-wallet/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	d := dependencyInject()

	r := gin.Default()

	r.GET("/health", d.HealthcheckAPI.HealthcheckHandlerHTTP)

	walletV1 := r.Group("/wallet/v1")
	walletV1.POST("/", d.WalletAPI.Create)
	walletV1.PUT("/balance/credit", d.MiddlewareValidateToken, d.WalletAPI.CreditBalance)
	walletV1.PUT("/balance/debit", d.MiddlewareValidateToken, d.WalletAPI.DebitBalance)
	walletV1.GET("/balance", d.MiddlewareValidateToken, d.WalletAPI.GetBalance)
	walletV1.GET("/history", d.MiddlewareValidateToken, d.WalletAPI.GetWalletHistory)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	HealthcheckAPI interfaces.IHealthcheckHandler
	WalletAPI      interfaces.IWalletHandler
	External       interfaces.IExternal
}

func dependencyInject() Dependency {
	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	walletRepo := &repository.WalletRepo{
		DB: helpers.DB,
	}

	walletService := &services.WalletService{
		WalletRepo: walletRepo,
	}

	walletApi := &api.WalletHandler{
		WalletService: walletService,
	}

	external := &external.External{}

	return Dependency{
		HealthcheckAPI: healthcheckAPI,
		WalletAPI:      walletApi,
		External:       external,
	}

}
