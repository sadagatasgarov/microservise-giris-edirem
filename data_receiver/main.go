package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sadagatasgarov/toll-calc/types"
)

func main() {
	recv, err := NewDataReceiver()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/ws", recv.handleWs)
	http.ListenAndServe(":30000", nil)
}

type DataReceiver struct {
	msgch chan types.OBUData
	conn  *websocket.Conn
	//prod  *kafka.Producer
	prod DataProducer
}

func NewDataReceiver() (*DataReceiver, error) {
	var (
		p DataProducer
		err error
	)

	p, err = NewKafkaProducer("topic")
	p = NewLogMiddleware(p)
	if err != nil {
		return nil, err
	}
	return &DataReceiver{
		msgch: make(chan types.OBUData, 128),
		prod:  p,
	}, nil
}

func (dr *DataReceiver) produceData(data types.OBUData) error {
	return dr.prod.ProduceData(data)
}

func (dr *DataReceiver) handleWs(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.conn = conn
	go dr.wsReceiveLoop()
}

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("New OBU connected client connected")
	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("Read error:", err)
			continue
		}

		if err := dr.produceData(data); err != nil {
			fmt.Println("Kafka rod eroru oldu", err)
		}

		//fmt.Printf("Received OBU data from [%d]--> Lat:%2.f, Long:%2.f \n", data.OBUID, data.Lat, data.Long)
		//dr.msgch <- data
	}
}
