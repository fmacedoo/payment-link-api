package main

import (
	"golink/handlers"
	"golink/shared/database"
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

	// Initialize the database
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		if err := database.CloseDB(); err != nil {
			log.Printf("Failed to close database: %v", err)
		}
	}()

	// Create Gin router
	router := gin.Default()

	// Define routes
	router.GET("/", handlers.Query)
	router.GET("/create-payment-link", handlers.CreatePaymentLinkHandler)
	router.POST("/webhook", handlers.StripeWebhookHandler)

	// Start the server
	router.Run(":9800")
}
