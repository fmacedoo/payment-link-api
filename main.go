package main

import (
	"golink/handlers"
	"golink/stripeapi"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the Stripe secret key from environment variables
	stripeSecretKey := os.Getenv("STRIPE_SECRET_KEY")
	if stripeSecretKey == "" {
		log.Fatalf("STRIPE_SECRET_KEY not set in .env file")
	}

	// Initialize Stripe
	stripeapi.InitializeStripe(stripeSecretKey)

	// Create Gin router
	router := gin.Default()

	// Define routes
	router.GET("/", handlers.Health)
	router.GET("/create-payment-link", handlers.CreatePaymentLinkHandler)
	router.POST("/webhook", handlers.StripeWebhookHandler)

	// Start the server
	router.Run(":9800")
}
