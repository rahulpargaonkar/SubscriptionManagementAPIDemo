package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Subcategory []Subcategory
}

type CategoryDTO struct  {
	Name string `json:"categoryName" binding:"required"`
}
