package domain

type Product struct {
	ProductID   int    `json:"product_id"`
	ProductCode string `json:"product_code"`
	ProductName string `json:"product_name"`
	Stock       int    `json:"stock"`
}
