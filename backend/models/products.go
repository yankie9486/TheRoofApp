package models

type Product struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	SlopeTypes  []string `json:"slopetypes"`
	Brand       string   `json:"brand"`
	Cost        float64  `json:"cost"`
	Price       float64  `json:"price"`
}
