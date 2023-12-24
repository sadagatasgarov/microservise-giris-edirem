package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	consumer  *kafka.Consumer
	isRunnung bool
}

func NewKafkaConsumer(topic string) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}
	c.SubscribeTopics([]string{topic}, nil)
	//c.Close()
	return &KafkaConsumer{
		consumer: c,
	}, nil
}

func (c *KafkaConsumer) Start() {
	logrus.Info("kafka transport started")
	c.readMessageLoop()
}

func (c *KafkaConsumer) readMessageLoop() {
	for c.isRunnung {
		msg, err := c.consumer.ReadMessage(time.Second)
		if err != nil {
			logrus.Errorf("Kafka consume error: %s", err)
			continue
		}
		fmt.Println(msg)
	}
}
