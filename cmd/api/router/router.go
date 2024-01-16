package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/talgat-ruby/planets-fact-site-api/cmd/api/handler"
	"github.com/talgat-ruby/planets-fact-site-api/cmd/api/middleware"
	apiT "github.com/talgat-ruby/planets-fact-site-api/cmd/api/types"
	dbT "github.com/talgat-ruby/planets-fact-site-api/cmd/db/types"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App, api apiT.Api, db dbT.DB) {
	m := middleware.New(api, db)
	h := handler.New(api, db)

	m.Logger(app)
	m.Swagger(app)

	group := app.Group("/api")
	v1(group, h)
}
