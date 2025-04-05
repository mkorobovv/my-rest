package config

import (
	"time"

	"github.com/mkorobovv/my-rest/internal/app/infrastructure/kafka"
	"github.com/mkorobovv/my-rest/internal/app/infrastructure/postgres"
)

type Config struct {
	Application    Application
	Infrastructure Infrastructure
	Adapters       Adapters
}

type Application struct {
	Name    string
	Version string
}

type Infrastructure struct {
	OrdersDB       postgres.Config
	OrdersConsumer kafka.Config
	OrdersProducer kafka.Config
}

type Adapters struct {
	Primary   Primary
	Secondary Secondary
}

type Primary struct {
	HttpAdapter          HttpAdapter
	PprofAdapter         PprofAdapter
	KafkaAdapterConsumer KafkaAdapterConsumer
}

type Secondary struct {
	KafkaAdapterProducer KafkaAdapterProducer
}

type HttpAdapter struct {
	Server Server
	Router Router
}

type Router struct {
	Shutdown Shutdown
	Timeout  Timeout
}

type Shutdown struct {
	Duration time.Duration
}

type Timeout struct {
	Duration time.Duration
}

type PprofAdapter struct {
	Server Server
	Router Router
}

type Server struct {
	Port              string
	Name              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	ReadHeaderTimeout time.Duration
	ShutdownTimeout   time.Duration
}

type KafkaAdapterProducer struct {
	OrdersProducer kafka.ProducerConfig
}

type KafkaAdapterConsumer struct {
	OrdersConsumer kafka.ConsumerConfig
}
