package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahulpargaonkar/SubscriptionManagementAPIDemo/dtos"
	"github.com/rahulpargaonkar/SubscriptionManagementAPIDemo/models"
)

func CategoryRoutes(router *gin.RouterGroup) {
	router.POST("/addCategory", AddCategory)
	router.PUT("/updateCategory/:categoryName", UpdateCategory)
}

func AddCategory(c *gin.Context) {
	var json dtos.CategoryDTO
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var err = models.SaveRecord(&models.Category{Name: json.Name})
	if err != nil {

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "Category Created Successfully")

}

func UpdateCategory(c *gin.Context) {
	CategoryName := c.Param("categoryName")

	var json dtos.CategoryDTO

	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	paramMap := make(map[string]string)
	paramMap["name"] = CategoryName
	var err = models.UpdateRecord(&models.Category{Name: json.Name}, paramMap)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnavailableForLegalReasons, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, "Category Name Updated Successfully")
	}

}
