package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

const (
	broker1Address = "localhost:8080"
	topic          = "kafka_learn"
	groupID        = "gid"
)

func StartKafka() {
	configDetail := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{mbroker1Address},
		Topic:    topic,
		GroupID:  groupID,
		MaxBytes: 100,
	})
	reader := kafka.NewReader(configDetail)
	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("err====+>", err.Error())
		}
		fmt.Println("message ===== +>", message)
	}
}
