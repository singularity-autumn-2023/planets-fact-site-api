package main

import (
	"context"
	"time"

	"github.com/talgat-ruby/planets-fact-site-api/cmd/api"
	"github.com/talgat-ruby/planets-fact-site-api/cmd/db"
	"github.com/talgat-ruby/planets-fact-site-api/configs"
	_ "github.com/talgat-ruby/planets-fact-site-api/docs"
	"github.com/talgat-ruby/planets-fact-site-api/internal/constant"
	"github.com/talgat-ruby/planets-fact-site-api/pkg/logger"
)

const idleTimeout = 5 * time.Second

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// conf
	conf, err := configs.NewConfig(ctx)
	if err != nil {
		panic(err)
	}

	// setup logger
	log := logger.New(conf)

	// configure db service
	d, err := db.New(log.With("service", constant.DB), conf.DB)
	if err != nil {
		log.ErrorContext(
			ctx,
			"initialize service error",
			"service", "database",
			"error", err,
		)
		panic(err)
	}
	log.InfoContext(ctx, "initialize service", "service", "database")

	// configure gateway service
	srv := api.New(log.With("service", constant.Api), conf.Api)
	log.InfoContext(ctx, "initialize service", "service", "api")
	// start gateway service
	srv.Start(ctx, cancel, d)

	<-ctx.Done()
	// Your cleanup tasks go here
	log.InfoContext(ctx, "cleaning up ...")

	log.InfoContext(ctx, "server was successful shutdown.")
}
