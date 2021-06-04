package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahulpargaonkar/SubscriptionManagementAPIDemo/dtos"
	"github.com/rahulpargaonkar/SubscriptionManagementAPIDemo/models"
	"github.com/thoas/go-funk"
)

func ProductRoutes(router *gin.RouterGroup) {
	router.POST("/addProduct/:categoryName/:subcategoryName", AddProduct)
	router.PUT("/addProductsWeeklyPrices/:productName", UpdateProductWeeklyPrices)

}

func AddProduct(c *gin.Context) {
	var json dtos.ProductDTO
	SubcategoryName := c.Param("subcategoryName")
	CategoryName := c.Param("categoryName")
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	var category models.Category
	var subCategory models.Subcategory
	models.FetchTableRowId(&category, models.Category{Name: CategoryName})
	models.FetchTableRowId(&subCategory, models.Subcategory{Name: SubcategoryName})

	if category.ID == 0 || subCategory.ID == 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "Category or SubCategory does Not exist"})
	} else {
		var err = models.SaveRecord(&models.Product{Name: json.Name, Frequency: json.Frequency,
			CategoryID: int64(category.ID), SubcategoryID: subCategory.ID})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusCreated, "Product Added Successfully; Please add or update product weekly prices if not updated")
		}
	}

}

func UpdateProductWeeklyPrices(c *gin.Context) {
	productName := c.Param("productName")
	days := [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday"}
	var json []dtos.WeeklyProductPriceDTO
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	for element := range json {
		if !funk.Contains(days, json[element].Day) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "'" + json[element].Day + "' Not Valid"})

		}

	}

	var product models.Product
	models.FetchTableRowId(&product, models.Product{Name: productName})

	if product.ID == 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "'" + productName + "' Not exist"})

	} else {

		columnNames := []string{"product_id", "day"}

		for element := range json {

			paramMap := make(map[string]interface{})
			paramMap["product_id"] = product.ID
			paramMap["day"] = json[element].Day
			if err := models.Upsert(&models.WeeklyProductPrice{Day: json[element].Day,
				Price: json[element].Price, ProductID: int64(product.ID)}, columnNames, paramMap); err != nil {
				c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			}
			paramMap = nil
		}

	}
}
