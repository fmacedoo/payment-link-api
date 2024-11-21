package dto

import "golink/shared/models"

type PaymentLinkDTO struct {
	Id            string `json:"id"`
	PaymentLinkID string `json:"payment_link_id"`
	Url           string `json:"url"`
	Paid          bool   `json:"paid"`
}

func NewPaymentLinkDTO(id, paymentLinkId, url string) *PaymentLinkDTO {
	p := PaymentLinkDTO{Id: id, PaymentLinkID: paymentLinkId, Url: url}
	return &p
}

func ConvertPaymentLinkModelToDTO(model models.PaymentLinkModel) PaymentLinkDTO {
	return PaymentLinkDTO{
		Id:            model.ID,
		PaymentLinkID: model.ID, // Assuming this is equivalent to ID; adjust if needed
		Url:           model.Url,
		Paid:          model.Paid,
	}
}

func ConvertPaymentLinkModelsToDTOs(models []models.PaymentLinkModel) []PaymentLinkDTO {
	dtos := make([]PaymentLinkDTO, len(models))
	for i, model := range models {
		dtos[i] = ConvertPaymentLinkModelToDTO(model)
	}
	return dtos
}
