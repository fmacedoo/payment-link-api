package handlers

import (
	"net/http"

	"golink/shared/database"
	"golink/shared/models"
	"golink/stripeapi"

	"github.com/gin-gonic/gin"
)

// CreatePaymentLinkHandler handles the payment link creation
func CreatePaymentLinkHandler(c *gin.Context) {
	// Call Stripe API to create the payment link
	paymentLink, err := stripeapi.CreatePaymentLink()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Save the payment link to the database
	linkModel := models.PaymentLinkModel{
		ID:            paymentLink.Id,
		PaymentLinkID: paymentLink.PaymentLinkID,
		Url:           paymentLink.Url,
	}

	err = database.SavePaymentLink(linkModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save payment link"})
		return
	}

	// Respond with the payment link
	c.JSON(http.StatusOK, gin.H{"payment_link": paymentLink.Url})
}
