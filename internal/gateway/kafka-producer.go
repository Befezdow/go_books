package gateway

import (
	"context"
	"github.com/segmentio/kafka-go"
)

import (
	"github.com/befezdow/go-books-rest-api/internal/shared"
)

type KafkaProducer struct {
	config *Config
	writer *kafka.Writer
}

func NewKafkaProducer(config *Config) *KafkaProducer {
	var writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{config.KafkaAddr},
		Topic:    config.KafkaTopic,
		Balancer: &kafka.LeastBytes{},
	})

	return &KafkaProducer{
		config: config,
		writer: writer,
	}
}

func (s *KafkaProducer) SendOneMessage(message shared.KafkaMessage) error {
	var err = s.writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   message.Key,
			Value: message.Value,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *KafkaProducer) SendManyMessages(messages []shared.KafkaMessage) error {
	var mappedMessages []kafka.Message
	for _, elem := range messages {
		mappedMessages = append(mappedMessages, kafka.Message{Key: elem.Key, Value: elem.Value})
	}

	var err = s.writer.WriteMessages(context.Background(),
		mappedMessages...,
	)
	if err != nil {
		return err
	}

	return nil
}
