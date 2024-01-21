package dto

import "fmt"

type ProductInputDTO struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  string  `json:"categoryId"`
	OwnerID     string  `json:"ownerId"`
}

func (p *ProductInputDTO) Validate() error {
	if p.Title == "" && p.Description == "" && p.Price <= 0 && p.CategoryID == "" && p.OwnerID == "" {
		return fmt.Errorf("request body is empty or malformed")
  }
	if p.Title == "" {
		return fmt.Errorf("title is required")
	}
	if p.Description == "" {
		return fmt.Errorf("description is required")
	}
	if p.Price <= 0 {
		return fmt.Errorf("price is required")
	}
	if p.CategoryID == "" {
		return fmt.Errorf("categoryId is required")
	}
	if p.OwnerID == "" {
		return fmt.Errorf("ownerId is required")
	}

	return nil
}

type ProductOutputDTO struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  string  `json:"categoryId"`
	OwnerID     string  `json:"ownerId"`
}
