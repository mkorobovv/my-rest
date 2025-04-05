package orders_consumer

import (
	"context"
	"errors"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/twmb/franz-go/pkg/kgo"
)

var ErrClientClosed = errors.New("client closed")

const kafkaPrefix = "KAFKA:"

func (c *OrdersConsumer) Start(ctx context.Context) error {
	defer c.client.Close()

	for {
		events := make([]cloudevents.Event, 0, c.config.BatchSize)
		records := make([]*kgo.Record, 0, c.config.BatchSize)

		fetches := c.client.PollRecords(ctx, c.config.BatchSize)

		if fetches.Empty() {
			continue
		}

		if fetches.IsClientClosed() {
			return ErrClientClosed
		}

		fetches.EachRecord(func(record *kgo.Record) {
			event, err := recordToEvent(record)
			if err != nil {
				c.logger.Error(kafkaPrefix, err.Error())

				return
			}

			events = append(events, event)
			records = append(records, record)
		})

		toCommit := make([]*kgo.Record, 0, len(records))

		for i, event := range events {
			err := c.handler.Process(ctx, event)
			if err != nil {
				c.logger.Error(kafkaPrefix, err.Error())

				break
			}

			toCommit = append(toCommit, records[i])
		}

		err := c.client.CommitRecords(ctx, toCommit...)
		if err != nil {
			c.logger.Error(kafkaPrefix, err.Error())
		}
	}
}
