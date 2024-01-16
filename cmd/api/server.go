package api

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"

	"github.com/talgat-ruby/planets-fact-site-api/cmd/api/router"
	apiT "github.com/talgat-ruby/planets-fact-site-api/cmd/api/types"
	dbT "github.com/talgat-ruby/planets-fact-site-api/cmd/db/types"
	"github.com/talgat-ruby/planets-fact-site-api/configs"
)

type server struct {
	log  *slog.Logger
	conf *configs.ApiConfig
	app  *fiber.App
}

func New(log *slog.Logger, conf *configs.ApiConfig) apiT.Api {
	s := &server{
		log:  log,
		conf: conf,
	}

	return s
}

func (s *server) GetLog() *slog.Logger {
	return s.log
}

func (s *server) Start(ctx context.Context, cancel context.CancelFunc, db dbT.DB) {
	app := fiber.New(fiber.Config{
		IdleTimeout: s.conf.IdleTimeout,
	})
	router.SetupRoutes(app, s, db)

	// Listen from s different goroutine
	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", s.conf.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.log.ErrorContext(ctx, "server error", "error", err)
		}

		cancel()
	}()

	shutdown := make(chan os.Signal, 1)                    // Create channel to signify s signal being sent
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM) //  When an interrupt or termination signal is sent, notify the channel

	go func() {
		_ = <-shutdown

		s.log.WarnContext(ctx, "gracefully shutting down...")
		if err := app.ShutdownWithContext(ctx); err != nil {
			s.log.ErrorContext(ctx, "server shutdown error", "error", err)
		}
	}()
}
