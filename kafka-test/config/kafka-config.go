package config

import "github.com/confluentinc/confluent-kafka-go/kafka"

func NewConsumer() *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	return c
}

func NewProducer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
	})
	if err != nil {
		panic(err)
	}
	return p
}
