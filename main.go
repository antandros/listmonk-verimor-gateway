package main

import (
	"encoding/json"

	"github.com/antandros/listmonk-verimor-gateway/config"
	"github.com/antandros/listmonk-verimor-gateway/handler"
	"github.com/antandros/listmonk-verimor-gateway/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ServerHeader: "Fiber",
		AppName:      "Listmonk-Verimor Gateway v1.0.1",
	})

	username := config.Config("USERNAME")
	password := config.Config("PASSWORD")

	app.Use(requestid.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			username: password,
		},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(418).JSON(&fiber.Map{"status": "I'm a teapot"})
	})

	app.Post("/notify", func(c *fiber.Ctx) error {
		payload := models.Listmonk{}
		json.Unmarshal(c.Body(), &payload)
		err := handler.SendSMS(payload)
		if err == nil {
			return c.JSON(fiber.Map{"status": "OK"})
		} else {
			return fiber.NewError(500, err.Error())
		}

	})

	app.Listen(config.Config("APP_BIND"))
}
