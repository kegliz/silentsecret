package server

import (
	"context"

	"kegnet.dev/silentsecret/internal/server/logger"
	"kegnet.dev/silentsecret/internal/server/router"
)

type (
	EngineOptions struct {
		Debug          bool
		TemplateFolder string
	}

	Server interface {
		Listen(port int) error
		Shutdown(ctx context.Context) error
	}
)

func NewLoggerAndRouter(options EngineOptions) (l *logger.Logger, r *router.Router) {
	l = logger.NewLogger(logger.LoggerOptions{
		Debug: options.Debug,
	})
	r = router.NewRouter(router.RouterOptions{
		Logger:         l,
		TemplateFolder: options.TemplateFolder,
	})
	return
}
