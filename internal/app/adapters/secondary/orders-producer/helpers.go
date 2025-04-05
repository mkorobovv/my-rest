package orders_producer

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/twmb/franz-go/pkg/kgo"
)

func eventToRecord(event cloudevents.Event) (record *kgo.Record) {
	ceId := kgo.RecordHeader{
		Key:   "ce_id",
		Value: bytes(event.ID()),
	}

	ceTime := kgo.RecordHeader{
		Key:   "ce_time",
		Value: bytes(event.Time().String()),
	}

	ceSource := kgo.RecordHeader{
		Key:   "ce_source",
		Value: bytes(event.Source()),
	}

	ceType := kgo.RecordHeader{
		Key:   "ce_type",
		Value: bytes(event.Type()),
	}

	ceSpecVersion := kgo.RecordHeader{
		Key:   "ce_specversion",
		Value: bytes(event.SpecVersion()),
	}

	contentType := kgo.RecordHeader{
		Key:   "content-type",
		Value: bytes(event.DataContentType()),
	}

	return &kgo.Record{
		Headers: []kgo.RecordHeader{
			ceId,
			ceTime,
			ceSource,
			ceType,
			ceSpecVersion,
			contentType,
		},
		Topic: event.Type(),
		Value: event.Data(),
	}
}

func bytes(s string) []byte {
	return []byte(s)
}
