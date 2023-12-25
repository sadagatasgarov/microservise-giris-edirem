package main

import "log"
const kafkaTopic = "topic"

func main() {

	service:=NewCalculatorService()
	
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, service)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
