package types

import (
	"context"

	"github.com/talgat-ruby/planets-fact-site-api/cmd/db/model"
)

type DB interface {
	GetPlanetsList(ctx context.Context) ([]model.PlanetItem, error)
	GetPlanet(ctx context.Context, id string) (*model.Planet, error)
}
