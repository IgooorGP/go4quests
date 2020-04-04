package config

import (
	"fmt"
	"github.com/IgooorGP/go4quests/go4quests/controllers"
	"github.com/IgooorGP/go4quests/go4quests/infra/persistence"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"os"
)

// Attempts to read a variable from env variables. If not possible, returns the default value.
func GetEnvValueOrUseDefault(envName string, envDefault string) string {
	envValue, fromEnv := os.LookupEnv(envName)

	if !fromEnv {
		return envDefault
	}

	return envValue
}

// Setups the logs of the application: level, format, etc.
func SetupLogs(loggingLevelRaw string, loggingLevelMap map[string]zerolog.Level) {
	loggingLevel, foundOnMap := loggingLevelMap[loggingLevelRaw]

	// logging level MUST be properly configured
	if !foundOnMap {
		panic(fmt.Errorf("unknown logging level: %s", loggingLevelRaw))
	}

	zerolog.SetGlobalLevel(loggingLevel)
}

// Sets env variables from .env.local file, if present
func LoadDotEnvIfAvailable(file string) error {
	err := godotenv.Load(file)

	if err != nil {
		return err
	}

	return nil
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
	ginMode := GinMode
	bindAddress := SocketBindAddress
	bindPort := SocketBindPort
	logLevel := LogLevel
	loggingLevelMap := LoggingLevelMap
	ginModeMap := GinModeMap

	// Setups Gin for prd ou dev
	SetupGin(ginMode, ginModeMap)

	// Setups logs
	SetupLogs(logLevel, loggingLevelMap)

	// Database setup
	_ = persistence.CreateDatabaseConnection(
		DatabaseEngine,
		DatabaseHost,
		DatabasePort,
		DatabaseName,
		DatabaseAppUser,
		DatabaseAppPassword,
		DatabaseUseSSL,
	)

	// Setups routes
	applicationRouter := gin.Default()
	SetupRoutes(applicationRouter)

	return applicationRouter, fmt.Sprintf("%s:%s", bindAddress, bindPort)
}
