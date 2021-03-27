package models

//Quote create a quote for an estimate
type Quote struct {
	ID          int64       `json:"id"`
	QuoteItems  []QuoteItem `json:"quoteItems"`
	TotalAmount float64     `json:"totalAmount"`
	TotalCost   float64     `json:"costAmount"`
}

type QuoteItem struct {
	ID           int64      `json:"id"`
	RoofCovering string     `json:"roofCovering"`
	RoofTypes    []RoofType `json:"rooftypes"`
	Products     []Product  `json:"products"`
}
