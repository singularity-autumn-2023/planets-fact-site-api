package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/talgat-ruby/planets-fact-site-api/cmd/api/types"
)

func v1(api fiber.Router, h types.Handler) {
	g := api.Group("/v1")

	g.Get("/planets", h.V1GetAllPlanets)
	g.Get("/planets/:id", h.V1GetPlanetByID)
}
