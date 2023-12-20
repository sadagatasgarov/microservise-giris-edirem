package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const wsEndPoint = "ws://localhost:30000/ws"

var sendInterval = time.Second

// func sebdOBUData(conn *websocket.Conn, data OBUData) error {
// 	return conn.WriteJSON(data)
// }

func genLatLong() (float64, float64) {
	return genCoord(), genCoord()
}

func genCoord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f
}

func main() {
	obuIDS := generateOBUIDS(20)
	// conn, _, err := websocket.DefaultDialer.Dial(wsEndPoint, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for {
	// 	for i := 0; i < len(obuIDS); i++ {
	// 		lat, long := genLatLong()
	// 		data := types.OBUData{
	// 			OBUID: obuIDS[i],
	// 			Lat:   lat,
	// 			Long:  long,
	// 		}
	// 		if err := conn.WriteJSON(data); err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		//fmt.Printf("%+v\n", data)
	// 	}

	// 	time.Sleep(sendInterval)
	// }
	//fmt.Println(rand.New(rand.NewSource(time.Now().UnixNano())))

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
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
	
	for {
	topic := "myTopic2"
	for _, word := range obuIDS {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(strconv.Itoa(word)),
		}, nil)
	}
	}
	// Wait for message deliveries before shutting down
	p.Flush(1 * 1000)

}

func generateOBUIDS(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}
