package entities

import (
	"fmt"
	"time"
)

type Category struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCategory(name string) (*Category, error) {
	category := &Category{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// business rules
	err := category.isValid()
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c *Category) isValid() error {
	if len(c.Name) < 5 {
		return fmt.Errorf("category name must be greater than 5, got %d", len(c.Name))
	}

	return nil
}
