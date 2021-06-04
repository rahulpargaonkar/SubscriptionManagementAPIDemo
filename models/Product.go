package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name               string
	Frequency          string
	WeeklyProductPrice []WeeklyProductPrice
	CategoryID         int64
	SubcategoryID      int64
}
