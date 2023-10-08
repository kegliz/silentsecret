package app

import (
	"context"

	"kegnet.dev/silentsecret/internal/config"
	"kegnet.dev/silentsecret/internal/server/logger"
	"kegnet.dev/silentsecret/internal/server/router"

	"kegnet.dev/silentsecret/internal/server"
)

type (
	ServerOptions struct {
		C       *config.Config
		Version string
	}

	appServer struct {
		logger  *logger.Logger
		router  *router.Router
		version string
	}

	appServerOptions struct {
		logger  *logger.Logger
		router  *router.Router
		version string
	}
)

// newAppServer creates a new appServer.
func newAppServer(options appServerOptions) *appServer {
	a := &appServer{
		logger:  options.logger,
		router:  options.router,
		version: options.version,
	}
	a.router.SetRoutes(a.routes())
	return a
}

// Listen implements server.Server.
func (a *appServer) Listen(port int) error {
	a.logger.Debug().Str("version", a.version).Msg("debug standup server")
	a.logger.Info(). //Str("host", "").
				Int("port", port).
				Msg("Starting silent secret service")
	return a.router.Start(port)
}

// Shutdown implements server.Server.
func (a *appServer) Shutdown(ctx context.Context) error {
	return a.router.Shutdown(ctx)
}

func NewServer(options ServerOptions) (server.Server, error) {
	l, r := server.NewLoggerAndRouter(server.EngineOptions{
		Debug:          options.C.GetBool("debug"),
		TemplateFolder: options.C.GetString("templatefolder"),
	})

	app := newAppServer(appServerOptions{
		logger:  l,
		router:  r,
		version: options.Version,
	})

	return app, nil
}
