package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	// Set up Kafka consumer configuration
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",         // broker addr
		"group.id":          "console-consumer-21564", // consumer group
		"auto.offset.reset": "earliest",               // earliest offset
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		fmt.Println("Failed to create New consumer")
		return
	}

	consumer.SubscribeTopics([]string{"testTopic"}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			data := string(msg.Value)
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, data)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
