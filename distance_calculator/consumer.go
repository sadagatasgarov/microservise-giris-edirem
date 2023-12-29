package main

import (
	"encoding/json"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sadagatasgarov/toll-calc/aggregator/client"
	"github.com/sadagatasgarov/toll-calc/types"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	consumer    *kafka.Consumer
	isRunnung   bool
	calcService CalculatorServicer
	aggClient   *client.HTTPClient
	
}

func NewKafkaConsumer(topic string, svc CalculatorServicer, aggClient *client.HTTPClient) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}
	c.SubscribeTopics([]string{topic}, nil)

	return &KafkaConsumer{
		consumer:    c,
		calcService: svc,
		aggClient:   aggClient,
	}, nil
}

func (c *KafkaConsumer) Start() {
	c.isRunnung = true
	logrus.Info("kafka transport started")
	c.readMessageLoop()
}

func (c *KafkaConsumer) readMessageLoop() {
	for c.isRunnung {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			logrus.Errorf("Kafka consume error: %s", err)
			continue
		}
		// if err == nil {
		// 	fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		// } else if !err.(kafka.Error).IsTimeout() {
		// 	// The client will automatically try to recover from all errors.
		// 	// Timeout is not considered an error because it is raised by
		// 	// ReadMessage in absence of messages.
		// 	fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		// }

		var data types.OBUData
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			logrus.Errorf("serialization error: %s", err)
			continue
		}
		distance, err := c.calcService.CalculateDistance(data)
		if err != nil {
			logrus.Errorf("calculation error: %s", err)
			continue
		}
		req := types.Distance{
			Value: distance,
			Unix:  time.Now().UnixNano(),
			OBUID: data.OBUID,
		}
		if err := c.aggClient.AggregateInvoice(req); err != nil {
			logrus.Error("agregate error: ", err)
			continue
		}
	}
}
