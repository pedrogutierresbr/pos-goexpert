package dto

type CreateProduct struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
