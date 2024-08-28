package product

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Product struct {
	ID        string
	Title     string
	Price     float64
	CreatedAt time.Time
}

func NewProduct(title string, price float64) (Product, error) {
	id, err := gonanoid.New()
	if err != nil {
		return Product{}, err
	}
	return Product{
		ID:        id,
		Title:     title,
		Price:     price,
		CreatedAt: time.Now(),
	}, nil
}
