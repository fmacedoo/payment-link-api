package stripeapi

import (
	"fmt"
	"golink/shared/constants"
	"golink/shared/models"
	"golink/shared/utils"
	"log"
	"os"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/paymentlink"
)

// InitializeStripe sets up Stripe with the secret key
func InitializeStripe(secretKey string) {
	stripe.Key = secretKey
}

// CreatePaymentLink creates a Stripe payment link
func CreatePaymentLink() (*models.GoLinkPaymentLink, error) {
	stripePriceId := os.Getenv("STRIPE_PRICE_ID")
	if stripePriceId == "" {
		log.Fatalf("STRIPE_WEBHOOK_SECRET not set in .env file")
	}

	id := utils.GenerateRandomString(12)

	params := &stripe.PaymentLinkParams{
		LineItems: []*stripe.PaymentLinkLineItemParams{
			{
				Price:    stripe.String(stripePriceId),
				Quantity: stripe.Int64(1),
			},
		},
		Metadata: map[string]string{
			constants.GOLINK_IDENTIFIER: id,
		},
	}

	link, err := paymentlink.New(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment link: %w", err)
	}

	return models.NewGoLinkPaymentLink(link.URL, id), nil
}