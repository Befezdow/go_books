package core

import "time"

type Config struct {
	BindAddr            string        `toml:"bind_addr"`
	LogLevel            string        `toml:"log_level"`
	KafkaAddr           string        `toml:"kafka_addr"`
	KafkaTopic          string        `toml:"kafka_topic"`
	KafkaPartition      int           `toml:"kafka_partition"`
	KafkaTimeout        time.Duration `toml:"kafka_timeout"`
	KafkaMaxMessageSize int           `toml:"kafka_max_message_size"`
	KafkaMinBatchSize   int           `toml:"kafka_min_batch_size"`
	KafkaMaxBatchSize   int           `toml:"kafka_max_batch_size"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:            ":8080",
		LogLevel:            "debug",
		KafkaAddr:           "localhost: 9092",
		KafkaTopic:          "books",
		KafkaPartition:      0,
		KafkaTimeout:        10,
		KafkaMaxMessageSize: 1024,
		KafkaMinBatchSize:   0,
		KafkaMaxBatchSize:   102400,
	}
}
