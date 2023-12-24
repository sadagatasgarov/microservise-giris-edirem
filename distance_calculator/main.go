package main

import "log"

const topic = "topic"

func main() {
	kafkaConsumer, err := NewKafkaConsumer(topic)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
