# Go Kafka Example

This repository demonstrates how to set up a simple Go project to produce and consume messages using Confluent Kafka.

## Features

- Kafka producer and consumer implemented in Go
- Basic example of sending and receiving messages

## Requirements

- Go 1.15 or higher
- Confluent Kafka
- Docker (optional but recommended for setting up Kafka)

## Getting Started

### Installation

1. **Clone the repository:**

    ```
    git clone https://github.com/muthukumar89uk/go-simple-kafka-example.git
    ```
 Click here to directly [download it](https://github.com/muthukumar89uk/go-simple-kafka-example/zipball/master). 

2. **Install Go dependencies:**

    ```
    go mod tidy
    ```
3. **Install Kafka locally:**

 Follow the instructions for your operating system on the [Official Confluent Kafka Website](https://docs.confluent.io/platform/current/platform-quickstart.html).

### Run the Application

1. **Run the Kafka producer:**

    ```
    go run producer.go
    ```

2. **Run the Kafka consumer:**

    ```
    go run consumer.go
    ```



