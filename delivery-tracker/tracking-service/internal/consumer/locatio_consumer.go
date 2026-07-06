package consumer

import (
	"context"
	// "log"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader 
}

func NewConsumer () *Consumer{
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic: "delivery-location",
		GroupID: "tracking-service",
	})

	return &Consumer{
		reader: reader,
	}
}

func (c *Consumer) Read() (kafka.Message,error){
	return c.reader.ReadMessage(context.Background())
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
