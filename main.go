package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rahulpargaonkar/SubscriptionManagementAPIDemo/controller"
	"github.com/rahulpargaonkar/SubscriptionManagementAPIDemo/models"
)

func main() {
	fmt.Println("Hello GO")
	database := models.InitDB()
	models.MigrateDB(database)

	router := gin.Default()
	apiRouteGroup := router.Group("/api")

	controller.CategoryRoutes(apiRouteGroup.Group("/categories"))
	controller.SubcategoryRoutes(apiRouteGroup.Group("/subCategory"))
	controller.ProductRoutes(apiRouteGroup.Group("/product"))
	controller.SubscriptionRoutes(apiRouteGroup.Group("/subscription"))
	controller.UserRoutes(apiRouteGroup.Group("/user"))
	router.Run(":8080")

}
