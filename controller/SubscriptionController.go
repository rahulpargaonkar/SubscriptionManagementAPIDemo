package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rahulpargaonkar/SubscriptionManagementAPIDemo/models"
)

func SubscriptionRoutes(router *gin.RouterGroup) {
	router.GET("/getSubScriptionAmount/:productName/:startDate/:subscriptionType", GetSubsciptionAmount)
	router.POST("/AddSubscription/:email/:productName/:subscriptionType/:startDate", AddSubscription)

}

func GetSubsciptionAmount(c *gin.Context) {
	productName := c.Param("productName")
	date := c.Param("startDate")
	subscriptionType := c.Param("subscriptionType")

	startDate, _ := time.Parse("02-01-2006", date)

	amt, _ := models.GetSubscriptionAmount(productName, subscriptionType, startDate)
	Amount := fmt.Sprintf("%.2f", amt)
	if amt != 0 {
		c.JSON(http.StatusOK, gin.H{"message": subscriptionType + " Subscription Amount for " + productName + " is " + Amount})
	} else {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": productName + " is Invalid"})
	}

}

func AddSubscription(c *gin.Context) {
	email := c.Param("email")
	productName := c.Param("productName")
	date := c.Param("startDate")
	subscriptionType := c.Param("subscriptionType")
	startDate, _ := time.Parse("02-01-2006", date)
	amt, endDate := models.GetSubscriptionAmount(productName, subscriptionType, startDate)

	var product models.Product
	models.FetchTableRowId(&product, models.Product{Name: productName})

	var user models.User
	models.FetchTableRowId(&user, models.User{Email: email})

	if err := models.SaveRecord(&models.Subscription{Type: subscriptionType,
		StartDate: startDate, EndDate: endDate, ProductID: int64(product.ID), UserID: int64(user.ID), Amount: amt}); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, "Subscription created Sucessfully")
	}

}
