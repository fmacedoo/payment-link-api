package handlers

import (
	"net/http"

	"golink/stripeapi"

	"github.com/gin-gonic/gin"
)

// CreatePaymentLinkHandler handles the payment link creation
func CreatePaymentLinkHandler(c *gin.Context) {
	// Call Stripe API to create the payment link
	plink, err := stripeapi.CreatePaymentLink()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the payment link
	c.JSON(http.StatusOK, gin.H{"data": plink})
}
