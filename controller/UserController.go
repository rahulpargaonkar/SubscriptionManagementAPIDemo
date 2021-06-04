package controller

import (
	
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahulpargaonkar/SubscriptionManagementAPIDemo/dtos"
	"github.com/rahulpargaonkar/SubscriptionManagementAPIDemo/models"
)

func UserRoutes(router *gin.RouterGroup) {
	router.POST("/addUser", AddUser)
	router.GET("/GetAllActiveSubscriptions/:email", GetAllActiveSubscriptions)

}

func AddUser(c *gin.Context) {
	var json dtos.UserDTO
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if err := models.SaveRecord(&models.User{Name: json.Name, Email: json.Email}); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, "User Created Successfully")
	}

}

func GetAllActiveSubscriptions(c *gin.Context) {
	email := c.Param("email")
	var user models.User
	models.FetchTableRowId(&user, models.User{Email: email})
	if user.ID != 0 {
		var subscriptionsDTO []dtos.SubscriptionDTO
		query := `select p.Name as ProductName, c.Name as CategoryName, sc.Name As SubcategoryName, 
				  sb.type  as  SubsctiptionType, sb.start_date as StartDate,sb.end_date as EndDate,
				  sb.amount  as Amount from 
				  subscriptions sb, products p, categories c, subCategories sc, users u where 
				  sb.user_id=u.id and
	   	          sb.product_id=p.id and
				  p.category_id=c.id and p.subcategory_id=sc.id
				  and u.email='` + email + "';"
		
		models.GetRecordsInStruct(query, &subscriptionsDTO)
		if len(subscriptionsDTO) != 0 {
			c.JSON(http.StatusOK, subscriptionsDTO)
		} else {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "No records found"})

		}

	} else {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": email + " is not valid or user not registered"})
	}
}
