package entity

type Product struct {
	ID          string `json:"_id"`
	Title       string
	Price       float64
	Description string
	OwnerID     string
	CategoryID  string
}
