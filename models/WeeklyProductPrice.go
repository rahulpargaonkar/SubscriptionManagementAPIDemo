package models

type WeeklyProductPrice struct {
	//gorm.Model
	ProductID int64  `gorm:"foriegnKey"`
	Day       string `gorm:"unique"`
	Price     float64
}
