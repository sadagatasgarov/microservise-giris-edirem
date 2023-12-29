package main

import (
	"context"
	"log"
	"time"

	"github.com/sadagatasgarov/toll-calc/aggregator/client"
	"github.com/sadagatasgarov/toll-calc/types"
)

func main() {
	c, err := client.NewGRPCClient(":3001")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := c.Aggregate(context.Background(), &types.AggregateRequest{
		ObuID: 1,
		Value: 58.55,
		Unix:  time.Now().UnixNano(),
	}); err != nil {
		log.Fatal(err)
	}
}
