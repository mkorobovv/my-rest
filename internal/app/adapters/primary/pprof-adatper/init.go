package pprof_adatper

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"github.com/mkorobovv/my-rest/internal/app/config"
	"golang.org/x/sync/errgroup"
	"log/slog"
	"net/http"
	"net/http/pprof"
)

type PprofAdapter struct {
	logger *slog.Logger
	server *http.Server
	config config.PprofAdapter
}

func New(logger *slog.Logger, config config.PprofAdapter) *PprofAdapter {
	router := newPprofRouter()

	server := &http.Server{
		Handler:           router,
		ReadTimeout:       config.Server.ReadTimeout,
		WriteTimeout:      config.Server.WriteTimeout,
		ReadHeaderTimeout: config.Server.ReadHeaderTimeout,
		Addr:              config.Server.Port,
	}

	return &PprofAdapter{
		logger: logger,
		server: server,
		config: config,
	}
}

func newPprofRouter() *mux.Router {
	router := mux.NewRouter()

	pprofSubrouter := router.PathPrefix("/debug/pprof").Subrouter()

	pprofSubrouter.HandleFunc("/", pprof.Index)
	pprofSubrouter.HandleFunc("/cmdline", pprof.Cmdline)
	pprofSubrouter.HandleFunc("/profile", pprof.Profile)
	pprofSubrouter.HandleFunc("/symbol", pprof.Symbol)
	pprofSubrouter.HandleFunc("/trace", pprof.Trace)

	return router
}

func (a PprofAdapter) Start(ctx context.Context) error {
	a.logger.Info(a.config.Server.Name + " started")

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
