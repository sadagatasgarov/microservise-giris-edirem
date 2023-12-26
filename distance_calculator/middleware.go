package main

import (
	"time"

	"github.com/sadagatasgarov/toll-calc/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next CalculatorServicer
}

func NewLogMiddleware(next CalculatorServicer) CalculatorServicer {
	return &LogMiddleware{
		next: next,
	}
}
// Burada implementde anlasilmazligim var geleckede bunu arassdiracagam
func (m *LogMiddleware) CalculateDistance(data types.OBUData) (float64, error) {
	dist, err := m.next.CalculateDistance(data)

	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
			"dist": dist,
		}).Info("calculate distance")
	}(time.Now())

	return dist, err
}
