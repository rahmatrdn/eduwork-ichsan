package configs

import (
	"os"
	"sesi_12/modules/product"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

func RouterConfig(app *fiber.App, db *gorm.DB) {
	app.Use(logger.New())
	app.Use(recover.New())

	product.ProductRouter(app, db)

	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
