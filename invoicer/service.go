package main

import (
	"fmt"

	"github.com/sadagatasgarov/toll-calc/types"
)

type Aggregator interface {
	//CalculateInvoice(types.OBUData)
	AggregateDistance(types.Distance) error
}

type Storer interface {
	Insert(types.Distance) error
}

type InvoiceAggregator struct {
	store Storer
}

func (i *InvoiceAggregator) AggregateDistance(distance types.Distance) error {
	fmt.Println("processing and inserting distance in the stoarage", distance)
	return i.store.Insert(distance)
}