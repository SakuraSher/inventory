package main

import (
	inventory "inventory/InventoryService"
	"log"

	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config := inventory.LoadConfig()
	db, err := gorm.Open(postgres.Open(config.DBConnectionString), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to the datbase %v", err)
	}

	db.AutoMigrate(&inventory.Item{})

	repo := inventory.NewInventoryRepository(db)
	service := inventory.newInventoryService(repo)
	handlers := inventory.NewInventoryHandlers(service)

	e := echo.New()

	e.GET("/items", handlers.ListItems)
	e.POST("items", handlers.CreateItem)
	e.GET("items/:id", handlers.GetItem)
	e.PUT("items/:id/quantity", handlers.UpdateItemQuantity)
	e.DELETE("items/:id", handlers.DeleteItem)

	e.Logger.Fatal(e.Start(":" + config.ServerPort))

}
