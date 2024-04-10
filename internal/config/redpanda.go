package config

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
)

type KafkaService struct {
	brokers []string
	topic   string
}

func NewKafkaService(brokers []string, topic string) *KafkaService {
	return &KafkaService{
		brokers: brokers,
		topic:   topic,
	}
}

func (ks *KafkaService) SendMessage(key string, value interface{}) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true

	valueBytes, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	producer, err := sarama.NewSyncProducer(ks.brokers, config)
	if err != nil {
		return fmt.Errorf("failed to create Kafka producer: %w", err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: ks.topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(valueBytes),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to send Kafka message: %w", err)
	}

	fmt.Printf("Sent message to partition %d at offset %d\n", partition, offset)
	return nil
}
