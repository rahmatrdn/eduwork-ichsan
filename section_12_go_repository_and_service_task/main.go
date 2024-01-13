package main

import (
	"sesi_12/configs"
	"sesi_12/modules/product"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	db, err := configs.DatabaseConfig()
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&product.Product{})

	app := fiber.New()

	configs.RouterConfig(app, db)
}
