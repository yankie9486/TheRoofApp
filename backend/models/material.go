package models

type Material struct {
	ID       int64 `json:"id"`
	Products []Product
}
