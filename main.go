// Copyright (C) 2023 Antandros Teknoloji A.Ş.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
		AppName:      "Listmonk-Verimor Gateway v1.0.2",
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
