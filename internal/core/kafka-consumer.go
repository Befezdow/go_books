package core

import (
	"context"
	"github.com/segmentio/kafka-go"
)

import (
	"github.com/befezdow/go-books-rest-api/internal/shared"
)

type KafkaConsumer struct {
	config *Config
	reader *kafka.Reader
}

func NewKafkaConsumer(config *Config) *KafkaConsumer {
	var reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{config.KafkaAddr},
		Topic:     config.KafkaTopic,
		Partition: config.KafkaPartition,
		MinBytes:  config.KafkaMinBatchSize,
		MaxBytes:  config.KafkaMaxBatchSize,
	})

	return &KafkaConsumer{
		config: config,
		reader: reader,
	}
}

func (s *KafkaConsumer) ReceiveAllMessages() []shared.KafkaMessage {
	var result []shared.KafkaMessage
	for {
		message, err := s.reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		result = append(result, shared.KafkaMessage{Key: message.Key, Value: message.Value})
	}

	return result
}

func (s *KafkaConsumer) ReceiveOneMessage() (shared.KafkaMessage, error) {
	var message, err = s.reader.ReadMessage(context.Background())
	if err != nil {
		return shared.KafkaMessage{}, err
	}
	return shared.KafkaMessage{Key: message.Key, Value: message.Value}, nil
}
