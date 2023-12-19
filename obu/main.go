package main

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sadagatasgarov/toll-calc/types"
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
	conn, _, err := websocket.DefaultDialer.Dial(wsEndPoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		for i := 0; i < len(obuIDS); i++ {
			lat, long := genLatLong()
			data := types.OBUData{
				OBUID: obuIDS[i],
				Lat:   lat,
				Long:  long,
			}
			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
			//fmt.Printf("%+v\n", data)
		}

		time.Sleep(sendInterval)
	}
	//fmt.Println(rand.New(rand.NewSource(time.Now().UnixNano())))

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
