package models

import "time"

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateProductInput struct {
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
}

type UpdateProductInput struct {
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
}

type DeleteProductInput struct {
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
}
