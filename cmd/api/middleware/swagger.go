package middleware

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

func (m *middlewareObject) Swagger(app *fiber.App) {
	swaggerCfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "/",
	}

	app.Use(swagger.New(swaggerCfg))
}
