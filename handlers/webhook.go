package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/webhook"

	"golink/shared/constants"
)

// StripeWebhookHandler processes Stripe webhook events
func StripeWebhookHandler(c *gin.Context) {
	// Read the body from Stripe
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	stripeWebhookSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")
	if stripeWebhookSecret == "" {
		log.Fatalf("STRIPE_WEBHOOK_SECRET not set in .env file")
	}

	// Verify the webhook signature
	signature := c.GetHeader("Stripe-Signature")
	event, err := webhook.ConstructEvent(payload, signature, stripeWebhookSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid webhook signature"})
		return
	}

	// Handle the event
	switch event.Type {
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		if err := json.Unmarshal(event.Data.Raw, &paymentIntent); err != nil {
			log.Printf("Failed to parse payment intent: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
			return
		}

		log.Printf("%w", event)
		log.Printf("%w", paymentIntent)

		if customIdentifier, exists := paymentIntent.Metadata[constants.GOLINK_IDENTIFIER]; exists {
			log.Printf("Custom Identifier: %s", customIdentifier)
		} else {
			log.Println("No custom identifier found in metadata")
		}
	case "payment_intent.failed":
		log.Println("Payment failed.")
	default:
		log.Println("Unhandled event type:", event.Type)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
