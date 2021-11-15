package model

type ProductList struct {
	Products []Product `json:"products"`
}

type Product struct {
	Sku      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
}
