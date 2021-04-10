package entity

import (
	"math/rand"
	"time"
)

// 产品实体类
type Product struct {
	ID        int
	Name      string
	Price     float64
	ImageName string
	imagePath string
}

func NewProduct(name string, price float64) *Product {
	rand.Seed(time.Now().UnixNano())
	return &Product{
		ID:    rand.Int(),
		Name:  name,
		Price: price,
	}
}
