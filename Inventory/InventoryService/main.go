package main

import (
	"log"

	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gor
)

func main() {
	config := LoadConfig()
	db, err := gorm.Open(postgres.Open(config.DBConnectionString), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to the datbase %v", err)
	}

	db.AutoMigrate(&Item{})

	repo := NewInventoryRepository(db)
	service := newInventoryService(repo)
	handlers := NewInventoryHandlers(service)

	e := echo.New()

	e.GET("/items", handlers.ListItems)
	e.POST("items", handlers.CreateItem)
	e.GET("items/:id", handlers.GetItem)
	e.PUT("items/:id/quantity", handlers.UpdateItemQuantity)
	e.DELETE("items/:id", handlers.DeleteItem)

	e.Logger.Fatal(e.Start(":" + config.ServerPort))

}
