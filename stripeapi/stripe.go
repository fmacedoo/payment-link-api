package stripeapi

import (
	"fmt"
	"golink/shared/constants"
	"golink/shared/dto"
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
func CreatePaymentLink() (*dto.PaymentLinkDTO, error) {
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
		PaymentIntentData: &stripe.PaymentLinkPaymentIntentDataParams{
			Metadata: map[string]string{
				constants.GOLINK_IDENTIFIER: id,
			},
		},
	}

	link, err := paymentlink.New(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment link: %w", err)
	}

	return dto.NewPaymentLinkDTO(id, link.ID, link.URL), nil
}

func DisablePaymentLink(id string) error {
	_, err := paymentlink.Update(id, &stripe.PaymentLinkParams{
		Active: stripe.Bool(false),
	})

	return err
}
