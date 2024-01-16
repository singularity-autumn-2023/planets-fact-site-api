package handler

import (
	"log/slog"

	"github.com/talgat-ruby/planets-fact-site-api/cmd/api/types"
	dbT "github.com/talgat-ruby/planets-fact-site-api/cmd/db/types"
)

type handlerObject struct {
	db  dbT.DB
	log *slog.Logger
}

func New(api types.Api, db dbT.DB) types.Handler {
	return &handlerObject{
		db:  db,
		log: api.GetLog(),
	}
}
