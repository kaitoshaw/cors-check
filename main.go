package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Email struct {
	Email string `json:"email" binding:"required,email"`
}

func main() {
	r := gin.Default()

	// Configure Cors
	r.Use(cors.New(cors.Config{
		// Allow specific origins
		AllowOrigins: []string{"http://localhost:3000"}, // Change to your frontend's origin
		// Allow all HTTP methods
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		// Allow all headers
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		// Allow sending cookies cross-domain
		AllowCredentials: true,
		// Set maximum age for preflight requests (OPTIONS)
		MaxAge: 12 * time.Hour,
	}))

	// Simple POST Request
	r.POST("/api/subscribe", func(c *gin.Context) {
		var email Email

		if err := c.ShouldBindJSON(&email); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("Received email:", email.Email)
		c.JSON(http.StatusOK, gin.H{"message": "Subscription successful!"})
	})

	// Run server
	r.Run(":9000")
}
