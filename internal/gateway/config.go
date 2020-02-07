package gateway

import "time"

type Config struct {
	BindAddr       string        `toml:"bind_addr"`
	LogLevel       string        `toml:"log_level"`
	KafkaAddr      string        `toml:"kafka_addr"`
	KafkaTopic     string        `toml:"kafka_topic"`
	KafkaPartition int           `toml:"kafka_partition"`
	KafkaTimeout   time.Duration `toml:"kafka_timeout"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:       ":8080",
		LogLevel:       "debug",
		KafkaAddr:      "localhost: 9092",
		KafkaTopic:     "books",
		KafkaPartition: 0,
		KafkaTimeout:   10,
	}
}
