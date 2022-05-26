package main

import (
	"fmt"
	"hacktiv8-assignment2/controllers"
	"hacktiv8-assignment2/repositories"
	"hacktiv8-assignment2/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "tasrifin"
	DB_PASSWORD = "tasrifin"
	DB_NAME     = "dbtest"
	APP_PORT    = ":8888"
)

func main() {
	db := connectDB()
	router := gin.Default()

	orderRepo := repositories.NewOrderRepo(db)
	orderService := services.NewOrderService(orderRepo)

	itemRepo := repositories.NewItemRepo(db)
	itemService := services.NewItemService(itemRepo)
	orderController := controllers.NewOrderController(orderService, itemService)

	router.POST("/orders", orderController.CreateNewOrder)
	router.GET("/orders", orderController.GetAllOrdersWithItems)
	router.DELETE("/orders/:orderId", orderController.DeleteOrder)
	router.PUT("/orders/:orderId", orderController.UpdateOrder)
	log.Println("Server running ar port : ", APP_PORT)
	router.Run(APP_PORT)
}

func connectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname =%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		panic(err)
	}

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	log.Default().Println("Connection to Database is Successfull")
	// db.AutoMigrate(models.Order{}, models.Item{})
	return db
}
