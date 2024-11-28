package main

import (
	"os"

	"golang-grubhub/database"

	middleware "golang-grubhub/middleware"
	routes "golang-grubhub/routes"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Logger())

	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(204)
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Restaurant API",
		})
	})

	routes.UserRoutes(router)

	router.Use(middleware.Authentication())
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
