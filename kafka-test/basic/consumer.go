package basic

import (
	"fmt"
	"kafka-test/config"
)

func Consumer() {
	c := config.NewConsumer()

	c.SubscribeTopics([]string{"weather"}, nil)
	defer c.Close()

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s : %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error : %v (%v)\n", err, msg)
		}
	}
}
