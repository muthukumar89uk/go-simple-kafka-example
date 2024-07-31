package main

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type data struct {
	Id      int
	Name    string
	Message string
}

func main() {
	topic := "testTopic"

	data := []data{
		{Id: 1, Name: "Steve", Message: "Welcome to Golang Team"},
		{Id: 2, Name: "Jack", Message: "Welcome to Golang Team"},
		{Id: 3, Name: "John", Message: "Welcome to Golang Team"},
		{Id: 4, Name: "Akshay", Message: "Welcome to Golang Team"},
	}

	dataJson, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Failed to Parse the data")
		return
	}
	
	Producer(string(dataJson), topic)
}

func Producer(message, topic string) {
	fmt.Println("Started the producer")

	// create Kafka producer
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		fmt.Println("Failed to create a New Produce")
		return
	}
	defer producer.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	// Close Kafka Producer
	producer.Flush(2 * 1000)
}
