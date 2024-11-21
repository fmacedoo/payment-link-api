package handlers

import (
	"fmt"
	"golink/shared/database"
	"golink/shared/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePaymentLinkHandler handles the payment link creation
func Query(c *gin.Context) {
	data, dbPaymentLinkGetAllErr := database.GetAllPaymentLinks()
	if dbPaymentLinkGetAllErr != nil {
		fmt.Errorf("Payment Link failed to be queried on database")
		http.Error(c.Writer, "Payment Link failed to be queried on database", http.StatusInternalServerError)
	}

	// Convert models to DTOs
	dtos := dto.ConvertPaymentLinkModelsToDTOs(data)

	// Respond with the converted DTOs
	c.JSON(http.StatusOK, dtos)
}
