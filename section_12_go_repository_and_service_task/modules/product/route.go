package product

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProductRouter(app *fiber.App, db *gorm.DB) {
	productRepository := NewProductRepository(db)
	productService := NewProductService(productRepository)
	productHandler := NewProductHandler(productService)

	app.Post("/products", productHandler.CreateProduct)
	app.Get("/products", productHandler.GetAllProduct)
	app.Get("/products/:id", productHandler.GetProductByID)
	app.Put("/products/:id", productHandler.UpdateProduct)
	app.Delete("/products/:id", productHandler.DeleteProduct)
}
