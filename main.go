package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("No .env found. Ignoring...")
	}

	router := gin.Default()

	router.GET("/healthcheck", func(ctx *gin.Context) {
		fmt.Println("hi...")

		ctx.JSON(http.StatusOK, gin.H{"health": os.Getenv("GIN_SERVER_TCP_SOCKET_BIND_ADDRESS")})
	})

	router.Run("0.0.0.0:8080")
}
