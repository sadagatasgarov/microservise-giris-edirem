package main

import "log"
const kafkaTopic = "topic"

func main() {
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
