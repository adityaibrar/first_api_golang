package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Price       int32  `json:"price"`
	Description string `json:"description"`
}

type ProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Price       int32  `json:"price" validate:"required"`
	Description string `json:"description" validate:"required"`
}
