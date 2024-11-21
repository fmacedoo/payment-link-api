package handlers

import (
	"encoding/json"
	"fmt"
	"golink/shared/constants"
	"golink/shared/database"
	"golink/stripeapi"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/v81/paymentintent"
	"github.com/stripe/stripe-go/webhook"
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
		var pi stripe.PaymentIntent
		if err := json.Unmarshal(event.Data.Raw, &pi); err != nil {
			http.Error(c.Writer, "Failed to parse payment intent", http.StatusBadRequest)
			return
		}

		fmt.Printf("PaymentIntent succeeded: %s\n", pi.ID)

		// Optional: Retrieve the PaymentIntent to fetch additional details
		intent, err := paymentintent.Get(pi.ID, nil)
		if err != nil {
			fmt.Printf("Failed to retrieve payment intent: %v\n", err)
			http.Error(c.Writer, "Failed to retrieve payment intent", http.StatusInternalServerError)
			return
		}

		if identifier, exists := intent.Metadata[constants.GOLINK_IDENTIFIER]; exists {
			fmt.Printf("Payment was made with identifier: %s\n", identifier)
			dbPaymentLink, getPaymentLinkErr := database.GetPaymentLinkById(identifier)
			if getPaymentLinkErr != nil {
				fmt.Errorf("No payment link found with identifier")
				return
			}

			dbPaymentLinkUpdateErr := database.UpdatePaymentLink(dbPaymentLink.ID, map[string]interface{}{"Paid": true})
			if dbPaymentLinkUpdateErr != nil {
				fmt.Errorf("Payment Link failed to be updated on database")
				http.Error(c.Writer, "Payment Link failed to be updated on database", http.StatusInternalServerError)
				return
			}

			disabledPaymentLinkErr := stripeapi.DisablePaymentLink(dbPaymentLink.PaymentLinkID)
			if disabledPaymentLinkErr != nil {
				fmt.Errorf("Payment Link failed to be disabled")
				http.Error(c.Writer, "Payment Link failed to be disabled", http.StatusInternalServerError)
				return
			}
		} else {
			fmt.Println("No payment link metadata found")
		}
	case "payment_intent.failed":
		log.Println("Payment failed.")
	default:
		log.Println("Unhandled event type:", event.Type)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
