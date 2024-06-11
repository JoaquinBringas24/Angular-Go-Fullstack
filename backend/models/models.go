package models

type Item struct {
	name        string  `json:"name"`
	description string  `json:"description"`
	terms       string  `json:"terms"`
	sku         string  `json:"sku"`
	price       float32 `json:"price"`
	images      string  `json:"images"`
}
