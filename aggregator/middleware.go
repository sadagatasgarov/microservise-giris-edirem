package main

import (
	"time"

	"github.com/sadagatasgarov/toll-calc/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next Aggregator
}

func NewLogMiddleware(next Aggregator) Aggregator {
	return &LogMiddleware{
		next: next,
	}
}

// AggregateDistance implements Aggregator.
func (m *LogMiddleware) AggregateDistance(distance types.Distance) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
		}).Info()
	}(time.Now())
	err = m.next.AggregateDistance(distance)
	return
}
