package main

import (
	"log"
	"transaction-service/handler"
	"transaction-service/orders"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/service_transaction?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	orderRepository := orders.NewRepository(db)
	orderService := orders.NewService(orderRepository)
	orderHandler := handler.NewOrderHandler(orderService)

	router := gin.Default()

	router.Use(cors.Default())
	router.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"POST", "GET", "PATCH", "DELETE", "HEAD"},
			AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		}))

	api := router.Group("/api/v1")

	api.POST("/orders", orderHandler.CreateNewOrder)
	api.GET("/orders", orderHandler.Index)
	api.GET("/orders/:id", orderHandler.Show)

	router.Run(":3003")

}
