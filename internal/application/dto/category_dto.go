package dto

import "fmt"

type CategoryInputDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	OwnerID     string `json:"owner_id"`
}

func (c *CategoryInputDTO) Validate() error {
	if c.Title == "" && c.Description == "" && c.OwnerID == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if c.Title == "" {
		return fmt.Errorf("title is required")
	}
	if c.Description == "" {
		return fmt.Errorf("description is required")
	}
	if c.OwnerID == "" {
		return fmt.Errorf("owner_id is required")
	}

	return nil
}

type CategoryOutputDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	OwnerID     string `json:"owner_id"`
}
