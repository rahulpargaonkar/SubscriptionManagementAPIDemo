package controller

import (
	"net/http"
	

	"github.com/gin-gonic/gin"
	"github.com/rahulpargaonkar/SubscriptionManagementAPIDemo/dtos"
	"github.com/rahulpargaonkar/SubscriptionManagementAPIDemo/models"
)

func SubcategoryRoutes(router *gin.RouterGroup) {
	router.POST("/addSubcategories", AddSubcategories)
	router.PUT("/updateSubcategory/:categoryName/:subcategoryName", UpdateSubcategory)
}

func AddSubcategories(c *gin.Context) {
	var json dtos.SubcategoryListDTO

	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var category models.Category
	models.FetchTableRowId(&category, models.Category{Name: json.Name})
	if category.ID == 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "'" + json.Name + "' Not exist"})
	} else {
		for element := range json.Subcategories {

			if err := models.SaveRecord(&models.Subcategory{Name: json.Subcategories[element].SubcategoryName, CategoryID: int64(category.ID)}); err != nil {
				c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			}
		}
	}

}

func UpdateSubcategory(c *gin.Context) {
	SubcategoryName := c.Param("subcategoryName")
	CategoryName := c.Param("categoryName")
	var json dtos.SubcategoryDTO

	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	paramMap := make(map[string]string)
	paramMap["name"] = SubcategoryName
	var category models.Category
	models.FetchTableRowId(&category, models.Category{Name: CategoryName})
	var err = models.UpdateRecord(&models.Subcategory{Name: json.SubcategoryName, CategoryID: int64(category.ID)}, paramMap)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnavailableForLegalReasons, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, "SubCategory Name Updated Successfully")
	}

}
