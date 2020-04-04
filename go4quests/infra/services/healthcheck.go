package services

import (
	"github.com/rs/zerolog/log"
)

// Service interface
type HealthCheckServiceInterface interface {
	AreExternalServicesHealthy() map[string]bool
}

// Service type which implements the interface
type HealthCheckService struct{}

// Service interface implementation
func (healthCheckApp *HealthCheckService) AreExternalServicesHealthy() map[string]bool {
	log.Info().Msg("Checking external services...")
	healthCheckResult := map[string]bool{
		"isHealthy": true,
	}

	return healthCheckResult
}

// Factory function for the health check service
func NewHealthCheckService() HealthCheckServiceInterface {
	healthCheckService := &HealthCheckService{}

	return healthCheckService
}
