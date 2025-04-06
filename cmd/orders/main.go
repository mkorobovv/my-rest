package main

import (
	"context"
	http_adapter "github.com/mkorobovv/my-rest/internal/app/adapters/primary/http-adapter"
	orders_consumer "github.com/mkorobovv/my-rest/internal/app/adapters/primary/orders-consumer"
	os_singnal_adapter "github.com/mkorobovv/my-rest/internal/app/adapters/primary/os-singnal-adapter"
	pprof_adatper "github.com/mkorobovv/my-rest/internal/app/adapters/primary/pprof-adatper"
	orders_producer "github.com/mkorobovv/my-rest/internal/app/adapters/secondary/orders-producer"
	orders_repository "github.com/mkorobovv/my-rest/internal/app/adapters/secondary/repositories/orders-repository"
	api_service "github.com/mkorobovv/my-rest/internal/app/application/api-service"
	orders_processor "github.com/mkorobovv/my-rest/internal/app/application/orders-processor"
	"github.com/mkorobovv/my-rest/internal/app/infrastructure/kafka"
	"github.com/mkorobovv/my-rest/internal/app/infrastructure/postgres"
	"golang.org/x/sync/errgroup"
	"log"
	"log/slog"
	"os"

	"github.com/mkorobovv/my-rest/internal/app/config"
	"github.com/mkorobovv/my-rest/internal/pkg/logger"
)

func main() {
	cfg := config.New()

	l := logger.New()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := newApp(l, cfg)

	err := app.start(
		ctx,
		app.httpAdapter,
		app.ordersConsumer,
		app.ordersProcessor,
		app.pprofAdapter,
	)
	if err != nil {
		l.Error(err.Error(), "main")

		os.Exit(1)
	}
}

type App struct {
	httpAdapter     *http_adapter.HttpAdapter
	pprofAdapter    *pprof_adatper.PprofAdapter
	ordersConsumer  *orders_consumer.OrdersConsumer
	ordersProcessor *orders_processor.OrdersProcessor
	osSignalAdapter *os_singnal_adapter.OsSignalAdapter
}

func newApp(l *slog.Logger, cfg config.Config) *App {
	ordersDB := postgres.New(l, cfg.Infrastructure.OrdersDB)

	ordersConsumerConnection := kafka.New(cfg.Infrastructure.OrdersConsumer, cfg.Adapters.Primary.KafkaAdapterConsumer.OrdersConsumer.GroupID)
	ordersProducerConnection := kafka.New(cfg.Infrastructure.OrdersProducer, "")

	ordersRepository := orders_repository.New(l, ordersDB)

	ordersProducer := orders_producer.New(l, cfg.Adapters.Secondary.KafkaAdapterProducer.OrdersProducer, ordersProducerConnection)

	apiService := api_service.New(l, ordersRepository)

	ordersProcessor := orders_processor.New(l, ordersProducer)
	httpAdapter := http_adapter.New(l, cfg.Adapters.Primary.HttpAdapter, apiService)
	pprofAdapter := pprof_adatper.New(l, cfg.Adapters.Primary.PprofAdapter)
	ordersConsumer := orders_consumer.New(l, ordersConsumerConnection, cfg.Adapters.Primary.KafkaAdapterConsumer.OrdersConsumer, apiService)

	return &App{
		httpAdapter:     httpAdapter,
		ordersConsumer:  ordersConsumer,
		ordersProcessor: ordersProcessor,
		osSignalAdapter: os_singnal_adapter.New(),
		pprofAdapter:    pprofAdapter,
	}
}

func (a *App) start(ctx context.Context, starters ...starter) error {
	g, ctx := errgroup.WithContext(ctx)

	for _, s := range starters {
		f := func() error {
			err := s.Start(ctx)
			if err != nil {
				log.Println(err)

				log.Println("starting graceful shutdown")

				return err
			}

			return nil
		}

		g.Go(f)
	}

	err := g.Wait()
	if err != nil {
		return err
	}

	return nil
}

type starter interface {
	Start(ctx context.Context) error
}
