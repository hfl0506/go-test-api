package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hfl0506/go-test-api/pkg/books"
	"github.com/hfl0506/go-test-api/pkg/common/config"
	"github.com/hfl0506/go-test-api/pkg/common/db"
)


func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()

	api := app.Group("/api")

	db := db.Init(c.DBUrl)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("All fine")
	})

	books.RegisterRoutes(api, db)

	app.Listen(c.Port)
}