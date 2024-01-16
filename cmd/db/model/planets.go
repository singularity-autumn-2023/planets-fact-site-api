package model

import (
	"context"
	"database/sql"
	"errors"
)

type PlanetItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (m *Model) GetPlanetsList(ctx context.Context) ([]PlanetItem, error) {
	m.log.InfoContext(ctx, "start GetPlanetsList")

	sqlStatement := `
		SELECT id, name
		FROM planet
	`

	rows, err := m.db.QueryContext(ctx, sqlStatement)
	if err != nil {
		m.log.ErrorContext(ctx, "fail GetPlanetsList", "error", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			m.log.ErrorContext(ctx, "error GetPlanetsList rows close", "error", err)
		}
	}(rows)

	planets := []PlanetItem{}
	for rows.Next() {
		p := PlanetItem{}
		err = rows.Scan(&p.ID, &p.Name)
		if err != nil {
			m.log.ErrorContext(ctx, "fail GetPlanetsList", "error", err)
			return nil, err
		}
		planets = append(planets, p)
	}

	m.log.InfoContext(ctx, "success GetPlanetsList")
	return planets, nil
}

type Planet struct {
	ID               string
	Name             string
	Rotation         int
	Revolution       int
	Radius           int
	Temperature      int
	OverviewContent  string
	OverviewSource   string
	StructureContent string
	StructureSource  string
	GeologyContent   string
	GeologySource    string
}

func (m *Model) GetPlanet(ctx context.Context, id string) (*Planet, error) {
	m.log.InfoContext(ctx, "start GetPlanet", "id", id)

	sqlStatement := `
		SELECT 
		    id,
		    name,
		    rotation,
		    revolution,
		    radius,
		    temperature,
		    overview_content,
		    overview_source,
		    structure_content,
		    structure_source,
		    geology_content,
		    geology_source
		FROM planet
		WHERE id=?
	`

	row := m.db.QueryRowContext(ctx, sqlStatement, id)
	if err := row.Err(); err != nil {
		m.log.ErrorContext(ctx, "fail GetPlanet", "error", err)
		return nil, err
	}

	p := Planet{}
	if err := row.Scan(
		&p.ID,
		&p.Name,
		&p.Rotation,
		&p.Revolution,
		&p.Radius,
		&p.Temperature,
		&p.OverviewContent,
		&p.OverviewSource,
		&p.StructureContent,
		&p.StructureSource,
		&p.GeologyContent,
		&p.GeologySource,
	); errors.Is(err, sql.ErrNoRows) {
		m.log.ErrorContext(ctx, "fail GetPlanet::no rows found", "error", err)
		return nil, err
	} else if err != nil {
		m.log.ErrorContext(ctx, "fail GetPlanet", "error", err)
		return nil, err
	}

	m.log.InfoContext(ctx, "success GetPlanet")
	return &p, nil
}
