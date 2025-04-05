package http_adapter

import (
	"context"
	"errors"
	"github.com/mkorobovv/my-rest/internal/app/adapters/primary/http-adapter/controller"
	"github.com/mkorobovv/my-rest/internal/app/adapters/primary/http-adapter/router"
	api_service "github.com/mkorobovv/my-rest/internal/app/application/api-service"
	"github.com/mkorobovv/my-rest/internal/app/config"
	"golang.org/x/sync/errgroup"
	"log/slog"
	"net/http"
)

type HttpAdapter struct {
	logger *slog.Logger
	server *http.Server
	config config.HttpAdapter
}

func New(logger *slog.Logger, config config.HttpAdapter, svc *api_service.ApiService) *HttpAdapter {
	router := newRouter(logger, svc)

	server := &http.Server{
		Handler:           router,
		ReadTimeout:       config.Server.ReadTimeout,
		WriteTimeout:      config.Server.WriteTimeout,
		ReadHeaderTimeout: config.Server.ReadHeaderTimeout,
		Addr:              config.Server.Port,
	}

	return &HttpAdapter{
		logger: logger,
		server: server,
		config: config,
	}
}

func newRouter(logger *slog.Logger, svc *api_service.ApiService) http.Handler {
	r := router.New()

	ctr := controller.New(logger, svc)

	r.AppendRoutes(ctr)

	router := r.Router()

	return router
}

func (a *HttpAdapter) Start(ctx context.Context) error {
	a.logger.Info(a.config.Server.Name + "started")

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), a.config.Server.ShutdownTimeout)
		defer cancel()

		err := a.server.Shutdown(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	g.Go(func() error {
		err := a.server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}

		return nil
	})

	err := g.Wait()
	if err != nil {
		return err
	}

	return nil
}
