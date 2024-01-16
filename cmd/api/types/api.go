package types

import (
	"context"
	"log/slog"

	"github.com/gofiber/fiber/v2"

	dbT "github.com/talgat-ruby/planets-fact-site-api/cmd/db/types"
)

type Api interface {
	Start(ctx context.Context, cancel context.CancelFunc, d dbT.DB)
	GetLog() *slog.Logger
}

type Middleware interface {
	Logger(app *fiber.App)
	Swagger(app *fiber.App)
}

type Handler interface {
	V1GetAllPlanets(c *fiber.Ctx) error
	V1GetPlanetByID(c *fiber.Ctx) error
}
