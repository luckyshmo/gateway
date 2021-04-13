package kafkaQueue

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/luckyshmo/gateway/models"
	"github.com/luckyshmo/gateway/models/sensor"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type KafkaStore struct {
	writer *kafka.Writer
}

func (ka *KafkaStore) WriteData(data ...sensor.Sensor) error {
	for i := 0; i < len(data); i++ {
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("Key-%d", i)),
			Value: []byte(fmt.Sprint(uuid.New())),
		}
		err := ka.writer.WriteMessages(context.Background(), msg)
		if err != nil {
			return errors.Wrap(err, "Error writing message to Kafka")
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}

func (ka *KafkaStore) WriteRawData(data ...models.RawData) error {
	logrus.Info("write to kafka RAW")
	return nil
}

func NewKafkaStore(kafkaURL, topic string) *KafkaStore {
	return &KafkaStore{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(kafkaURL),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}
