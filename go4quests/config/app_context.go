package config

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"os"
	"path"
)

// Attempts to read a variable from env variables. If not possible, returns the default value.
func GetEnvValueOrUseDefault(envName string, envDefault string) string {
	envValue, fromEnv := os.LookupEnv(envName)

	if !fromEnv {
		return envDefault
	}

	return envValue
}

// Configuration struct's
type DatabaseConfig struct {
	DatabaseEngine   string
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseUseSSL   string
}

// Project root path and .env file loading (if available)
var ProjectRootPath = path.Join(os.Getenv("GOPATH"), "src/github.com/IgooorGP/go4quests")
var _ = godotenv.Load(path.Join(ProjectRootPath, ".env")) // loads env variables from .env file

// Host and port of the app
var SocketBindAddress = GetEnvValueOrUseDefault("GIN_SERVER_TCP_SOCKET_BIND_ADDRESS", "0.0.0.0")
var SocketBindPort = GetEnvValueOrUseDefault("GIN_SERVER_TCP_SOCKET_BIND_PORT", "8080")

// Database configurations
var DBConfig = DatabaseConfig{
	DatabaseEngine:   GetEnvValueOrUseDefault("DB_ENGINE", "postgres"),
	DatabaseHost:     GetEnvValueOrUseDefault("DB_HOST", "localhost"),
	DatabasePort:     GetEnvValueOrUseDefault("DB_PORT", "5432"),
	DatabaseName:     GetEnvValueOrUseDefault("DB_NAME", ""),
	DatabaseUser:     GetEnvValueOrUseDefault("DB_USER", ""),
	DatabasePassword: GetEnvValueOrUseDefault("DB_PASSWORD", ""),
	DatabaseUseSSL:   GetEnvValueOrUseDefault("DB_SSL", "disable"),
}

// Gin mode config
var GinMode = GetEnvValueOrUseDefault("GIN_MODE", "debug")
var GinModeMap = map[string]string{
	"release": gin.ReleaseMode,
	"debug":   gin.DebugMode,
}

// Logging config
var LoggingLevelMap = map[string]zerolog.Level{
	"panic": zerolog.PanicLevel,
	"fatal": zerolog.FatalLevel,
	"warn":  zerolog.WarnLevel,
	"info":  zerolog.InfoLevel,
	"debug": zerolog.DebugLevel,
}
var LogLevel = GetEnvValueOrUseDefault("LOG_LEVEL", "info")
