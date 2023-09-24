package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"

	"service-1/config"
)

func NewKafkaConn(cfg *config.Config) (*kafka.Conn, error) {
	return kafka.DialContext(context.Background(), "tcp", cfg.Kafka.Brokers[0])
}
