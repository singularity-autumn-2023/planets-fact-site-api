package logger

import (
	"log/slog"
	"os"

	"github.com/talgat-ruby/planets-fact-site-api/configs"
	"github.com/talgat-ruby/planets-fact-site-api/internal/constant"
)

func New(conf *configs.Config) *slog.Logger {
	if conf.Env != constant.EnvironmentLocal {
		return slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}

	return slog.New(slog.NewTextHandler(os.Stdout, nil))
}
