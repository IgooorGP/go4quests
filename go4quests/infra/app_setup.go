package infra

import (
	"fmt"
	"github.com/IgooorGP/go4quests/go4quests/config"
	"github.com/IgooorGP/go4quests/go4quests/controllers"
	"github.com/IgooorGP/go4quests/go4quests/infra/persistence"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// Setups the logs of the application: level, format, etc.
func SetupLogs(loggingLevelRaw string, loggingLevelMap map[string]zerolog.Level) {
	loggingLevel, foundOnMap := loggingLevelMap[loggingLevelRaw]

	// logging level MUST be properly configured
	if !foundOnMap {
		panic(fmt.Errorf("unknown logging level: %s", loggingLevelRaw))
	}

	zerolog.SetGlobalLevel(loggingLevel)
}

// Setups the router's routes of the application
func SetupRoutes(router *gin.Engine) {
	router.GET("/healthcheck", controllers.HealthCheckController)
}

// Setups GIN for dev or prd release
func SetupGin(ginModeRaw string, ginModeMap map[string]string) {
	ginMode, foundOnMap := ginModeMap[ginModeRaw]

	// logging level MUST be properly configured
	if !foundOnMap {
		panic(fmt.Errorf("unknown Gin mode: %s", ginModeRaw))
	}

	gin.SetMode(ginMode)
}

// Bootstraps (setups) the main application.
func SetupApplication() (*gin.Engine, string) {
	// Global config variables access
	ginMode := config.GinMode
	bindAddress := config.SocketBindAddress
	bindPort := config.SocketBindPort
	logLevel := config.LogLevel
	loggingLevelMap := config.LoggingLevelMap
	ginModeMap := config.GinModeMap

	// Setups Gin for prd ou dev
	SetupGin(ginMode, ginModeMap)

	// Setups logs
	SetupLogs(logLevel, loggingLevelMap)

	// Database setup
	_ = persistence.NewDatabase(config.DBConfig)

	// Setups routes
	applicationRouter := gin.Default()
	SetupRoutes(applicationRouter)

	return applicationRouter, fmt.Sprintf("%s:%s", bindAddress, bindPort)
}
