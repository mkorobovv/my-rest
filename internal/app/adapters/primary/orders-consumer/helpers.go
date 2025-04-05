package orders_consumer

import (
	"strings"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/twmb/franz-go/pkg/kgo"
)

func recordToEvent(record *kgo.Record) (cloudevents.Event, error) {
	event := cloudevents.NewEvent()

	err := event.SetData(cloudevents.ApplicationJSON, record.Value)
	if err != nil {
		return cloudevents.Event{}, err
	}

	headers := extractKeys(record)

	ceTime, ok := headers["ce_time"]
	if !ok {
		ceTime = time.Now().Format(time.DateTime)
	}

	ceTime = strings.Split(ceTime, ".")[0]

	t, err := time.Parse(time.DateTime, ceTime)
	if err != nil {
		return cloudevents.Event{}, err
	}

	event.SetTime(t)

	return event, nil
}

func extractKeys(record *kgo.Record) map[string]string {
	headers := make(map[string]string)

	for _, header := range record.Headers {
		headers[header.Key] = string(header.Value)
	}

	return headers
}
