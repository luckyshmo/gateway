package kafkaQueue

import (
	"github.com/luckyshmo/gateway/models"
	"github.com/segmentio/kafka-go"
)

type KafkaStore struct {
	writer *kafka.Writer
}

func (ka *KafkaStore) WriteData(...models.RawData) error {
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

// func mainProducer(address string, topic string) {

// 	writer := newKafkaWriter(address, topic)
// 	defer writer.Close()
// 	InfoLogger.Println("start producing ... !!")
// 	for i := 0; ; i++ {
// 		msg := kafka.Message{
// 			Key:   []byte(fmt.Sprintf("Key-%d", i)),
// 			Value: []byte(fmt.Sprint(uuid.New())),
// 		}
// 		err := writer.WriteMessages(context.Background(), msg)
// 		if err != nil {
// 			WarningLogger.Println(err)
// 		}
// 		time.Sleep(1 * time.Second)
// 	}
// }
