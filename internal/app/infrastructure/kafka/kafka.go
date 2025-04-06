package kafka

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/plain"
	"github.com/twmb/franz-go/pkg/sasl/scram"
	"github.com/twmb/franz-go/plugin/kslog"
)

const KafkaPrefix = "KAFKA:"

var ErrClientClosed = errors.New("client closed")

const (
	SASLMechanismPlain       = "PLAIN"
	SASLMechanismSCRAMSHA256 = "SCRAM-SHA-256"
	SASLMechanismSCRAMSHA512 = "SCRAM-SHA-512"
)

type Config struct {
	Hosts         string `config:"envVar"`
	User          string `config:"envVar"`
	Password      string `config:"envVar"`
	SASLMechanism string
}

type ProducerConfig struct {
	Topic string
}

type ConsumerConfig struct {
	Topic     string
	GroupID   string
	BatchSize int
}

func New(config Config, groupID string) *kgo.Client {
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn})
	logger := slog.
		New(h).
		With(slog.String("component", "KAFKA")).
		With(slog.String("group_id", groupID))

	seeds := strings.Split(config.Hosts, ",")

	options := []kgo.Opt{
		kgo.SeedBrokers(seeds...),
		kgo.WithLogger(kslog.New(logger)),
		kgo.RetryTimeout(30 * time.Second),
		kgo.RequiredAcks(kgo.AllISRAcks()),
	}

	//saslOption, err := getSASLOption(config)
	//if err != nil {
	//	panic(err)
	//}
	//
	//options = append(options, saslOption)

	if groupID != "" {
		options = append(options,
			kgo.ConsumerGroup(groupID),
			kgo.DisableAutoCommit(),
		)
	}

	client, err := kgo.NewClient(options...)
	if err != nil {
		errMsg := fmt.Sprintf("Can't connect: %v.\nMake sure a Kafka Server is running at: %s", err, config.Hosts)

		panic(errMsg)
	}

	return client
}

func getSASLOption(config Config) (kgo.Opt, error) {
	switch config.SASLMechanism {
	case SASLMechanismPlain:
		mechanism := plain.Auth{
			User: config.User,
			Pass: config.Password,
		}

		return kgo.SASL(mechanism.AsMechanism()), nil
	case SASLMechanismSCRAMSHA256:
		auth := scram.Auth{User: config.User, Pass: config.Password}

		return kgo.SASL(auth.AsSha256Mechanism()), nil
	case SASLMechanismSCRAMSHA512:
		auth := scram.Auth{User: config.User, Pass: config.Password}

		return kgo.SASL(auth.AsSha512Mechanism()), nil
	default:
		err := fmt.Errorf("unsupported SASL mechanism: %s", config.SASLMechanism)

		return nil, err
	}
}
