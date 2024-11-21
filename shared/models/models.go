package models

type GoLinkPaymentLink struct {
	Link string `json:"link"`
	Id   string `json:"id"`
}

func NewGoLinkPaymentLink(link string, id string) *GoLinkPaymentLink {
	p := GoLinkPaymentLink{Link: link, Id: id}
	return &p
}
