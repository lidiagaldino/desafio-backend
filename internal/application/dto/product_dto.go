package dto

type ProductInputDTO struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  string  `json:"categoryId"`
	OwnerID     string  `json:"ownerId"`
}

type ProductOutputDTO struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  string  `json:"categoryId"`
	OwnerID     string  `json:"ownerId"`
}
