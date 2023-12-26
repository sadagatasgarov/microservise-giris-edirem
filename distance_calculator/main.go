package main

import (
	"log"

	"github.com/sadagatasgarov/toll-calc/aggregator/client"
)

const (
	kafkaTopic         = "topic"
	aggregatorEndpoint = "http://127.0.0.1:3000/aggregate"
)

func main() {
	var (
		err error
		svc CalculatorServicer
	)
	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, client.NewClient(aggregatorEndpoint))
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
