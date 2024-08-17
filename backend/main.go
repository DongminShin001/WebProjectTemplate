package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })

    // Use "os" to get an environment variable
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if not set
    }

    r.Run(":" + port)
}
