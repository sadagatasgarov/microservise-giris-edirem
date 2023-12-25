package main

import (
	"fmt"

	"github.com/sadagatasgarov/toll-calc/types"
)

// we like the end our interface with (er) : shung
type CalculatorServicer interface {
	CalculateDistance(types.OBUData) (float64, error)
}

type CalculatorService struct {
}

func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

func (s *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	fmt.Println("calculating the distance")
	return 0.0, nil
}
