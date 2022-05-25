package basic

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"kafka-test/config"
	"kafka-test/sensor"
)

func Producer() {
	p := config.NewProducer()
	defer p.Close()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed : %s\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v, %s\n", ev.TopicPartition, ev.Value)
				}
			}
		}
	}()

	topic := "myTopic"
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		_ = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	topic2 := "weather"
	var humid [10]string
	humid = sensor.Humidity()

	for h, word := range humid {

		fmt.Printf("humid : %b\n", h)

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic2, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)

		p.Flush(15 * 1000)
	}

}
