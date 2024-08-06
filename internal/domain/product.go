package domain

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ProductID   uint   `gorm:"primaryKey"`
	ProductCode string `json:"product_code" gorm:"uniqueIndex;size:255" validate:"required,min=3,max=30"`
	ProductName string `json:"product_name" gorm:"size:255" validate:"required,min=3,max=100"`
	Stock       int    `json:"stock" validate:"gte=0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
