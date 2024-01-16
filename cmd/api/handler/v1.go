package handler

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/talgat-ruby/planets-fact-site-api/cmd/db/model"
	"github.com/talgat-ruby/planets-fact-site-api/pkg/random"
	"github.com/talgat-ruby/planets-fact-site-api/pkg/shuffle"
)

type V1GetAllPlanetsResponse struct {
	Count int                `json:"count"`
	Data  []model.PlanetItem `json:"data"`
}

// V1GetAllPlanets is a function to get planets data from database
// @Summary Get planets
// @Description Get random planets list from 1-8
// @Tags planets
// @Accept json
// @Produce json
// @Success 200 {object} V1GetAllPlanetsResponse{data=[]model.PlanetItem}
// @Failure 503 {object} V1GetAllPlanetsResponse{}
// @Router /api/v1/planets [get]
func (h *handlerObject) V1GetAllPlanets(c *fiber.Ctx) error {
	ctx := c.UserContext()
	h.log.InfoContext(ctx, "start V1GetAllPlanets")

	planets, err := h.db.GetPlanetsList(ctx)
	if err != nil {
		h.log.ErrorContext(
			ctx,
			"fail V1GetAllPlanets",
			"error", err,
		)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	randLen := random.MinMaxInt(1, 8)
	planets = shuffle.Slice(planets)[:randLen]

	h.log.InfoContext(ctx, "success V1GetAllPlanets")
	return c.JSON(V1GetAllPlanetsResponse{
		Count: len(planets),
		Data:  planets,
	})
}

type ItemInfo struct {
	Content string `json:"content"`
	Source  string `json:"source"`
}

type V1GetPlanetByIDResponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Rotation    int      `json:"rotation"`
	Revolution  int      `json:"revolution"`
	Radius      int      `json:"radius"`
	Temperature int      `json:"temperature"`
	Overview    ItemInfo `json:"overview"`
	Structure   ItemInfo `json:"structure"`
	Geology     ItemInfo `json:"geology"`
}

// V1GetPlanetByID is a function to get a planets by ID
// @Summary Get planet by ID
// @Description Get planet by ID
// @Tags planets
// @Accept json
// @Produce json
// @Param id path string true "Planet ID"
// @Success 200 {object} V1GetPlanetByIDResponse{}
// @Failure 404 {object} V1GetPlanetByIDResponse{}
// @Failure 503 {object} V1GetPlanetByIDResponse{}
// @Router /api/v1/planets/{id} [get]
func (h *handlerObject) V1GetPlanetByID(c *fiber.Ctx) error {
	ctx := c.UserContext()
	planetId := c.Params("id")
	h.log.InfoContext(ctx, "start V1GetPlanetByID", "planetId", planetId)

	planet, err := h.db.GetPlanet(ctx, planetId)
	if errors.Is(err, sql.ErrNoRows) {
		h.log.ErrorContext(
			ctx,
			"fail V1GetPlanetByID",
			"error::not found", err,
		)
		return c.Status(fiber.StatusNotFound).SendString(fmt.Sprintf("planet with id %s not found", planetId))
	} else if err != nil {
		h.log.ErrorContext(
			ctx,
			"fail V1GetPlanetByID",
			"error", err,
		)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	h.log.InfoContext(ctx, "success V1GetPlanetByID")
	return c.JSON(V1GetPlanetByIDResponse{
		ID:          planet.ID,
		Name:        planet.Name,
		Rotation:    planet.Rotation,
		Revolution:  planet.Revolution,
		Radius:      planet.Radius,
		Temperature: planet.Temperature,
		Overview: ItemInfo{
			Content: planet.OverviewContent,
			Source:  planet.OverviewSource,
		},
		Structure: ItemInfo{
			Content: planet.StructureContent,
			Source:  planet.StructureSource,
		},
		Geology: ItemInfo{
			Content: planet.GeologyContent,
			Source:  planet.GeologySource,
		},
	})
}
