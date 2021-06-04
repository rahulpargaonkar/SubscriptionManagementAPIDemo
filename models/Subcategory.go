package models

import (
	"gorm.io/gorm"
)

type Subcategory struct {
	gorm.Model
	ID         int64
	Name       string `gorm:"unique"`
	CategoryID int64
}
