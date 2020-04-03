package infra

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// Loads env vars from .env.local if available
// For production, k8s Config map or other env source should set these variables
var _ = LoadDotEnvIfAvailable(".env.local")

// Host and port of the app
var SocketBindAddress = GetEnvValueOrUseDefault("GIN_SERVER_TCP_SOCKET_BIND_ADDRESS", "info")
var SocketBindPort = GetEnvValueOrUseDefault("GIN_SERVER_TCP_SOCKET_BIND_PORT", "info")

// Gin mode configs
var GinMode = GetEnvValueOrUseDefault("GIN_MODE", "debug")
var GinModeMap = map[string]string{
	"release": gin.ReleaseMode,
	"debug":   gin.DebugMode,
}

// Logging configs
var LoggingLevelMap = map[string]zerolog.Level{
	"panic": zerolog.PanicLevel,
	"fatal": zerolog.FatalLevel,
	"warn":  zerolog.WarnLevel,
	"info":  zerolog.InfoLevel,
	"debug": zerolog.DebugLevel,
}
var LogLevel = GetEnvValueOrUseDefault("LOG_LEVEL", "info")
