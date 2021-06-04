package dtos

import "time"

type CategoryDTO struct {
	Name string `json:"categoryName" binding:"required"`
}

type SubcategoryListDTO struct {
	Name          string           `json:"categoryName" binding:"required"`
	Subcategories []SubcategoryDTO `json:"subCategories" binding:"required"`
}

type SubcategoryDTO struct {
	SubcategoryName string `json:"subcategoryName" binding:"required"`
}

type ProductDTO struct {
	Name      string `json:"productName" binding:"required"`
	Frequency string `json:"frequency" binding:"required"`
}

type WeeklyProductPriceDTO struct {
	Day   string  `json:"day" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

type ProductPricesDTO struct {
	Day   string
	Price float64
}

type UserDTO struct {
	Name  string `json:"userName" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type SubscriptionDTO struct {
	ProductName      string
	CategoryName     string
	SubcategoryName  string
	SubsctiptionType string
	StartDate        time.Time
	EndDate          time.Time
	Amount           float64
}
