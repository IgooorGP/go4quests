package controllers

import (
	"github.com/IgooorGP/go4quests/go4quests/infra"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Controller (interface)
func HealthCheckController(ctx *gin.Context) {
	healthCheckApp := infra.NewHealthCheckService()
	healthCheckResult := healthCheckApp.AreExternalServicesHealthy()

	ctx.JSON(http.StatusOK, healthCheckResult)
}
