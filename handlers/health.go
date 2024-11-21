package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePaymentLinkHandler handles the payment link creation
func Health(c *gin.Context) {
	// Respond with the payment link
	c.JSON(http.StatusOK, gin.H{"status": "healthy!"})
}
