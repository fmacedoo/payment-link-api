package models

type PaymentLinkModel struct {
	ID            string `gorm:"primaryKey"`
	PaymentLinkID string `gorm:"not null"`
	Url           string `gorm:"not null"`
	Paid          bool   `gorm:"not null;default:false"`
}

func NewPaymentLinkModel(id, paymentLinkId, url string) *PaymentLinkModel {
	p := PaymentLinkModel{ID: id, PaymentLinkID: paymentLinkId, Url: url}
	return &p
}
