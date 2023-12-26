package main

import "github.com/sadagatasgarov/toll-calc/types"

type LogMiddleware struct {
	next CalculatorServicer
}



func NewLogMiddleware(next CalculatorServicer) CalculatorServicer {
	return &LogMiddleware{
		next: next,
	}
}

func (m *LogMiddleware) CalculateDistance(data types.OBUData) (float64, error) {
	return 0.0, nil
}
