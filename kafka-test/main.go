package main

import (
	"fmt"
	"kafka-test/basic"
)

func main() {

	fmt.Println("Kafka Producer Example")
	go basic.Producer()

	fmt.Println("Kafka Consumer Example")
	go basic.Consumer()

	select {}
}
