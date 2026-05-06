package interfaces

import "github.com/gin-gonic/gin"

type IHealthcheckServices interface {
	HealthcheckServices() (string, error)
}

type IHealthcheckRepo interface {
}

type IHealthcheckHandler interface {
	HealthcheckHandlerHTTP(c *gin.Context)
}
