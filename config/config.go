package config

import (
	"time"

	"github.com/sony/gobreaker"
)

func NewCircuitBreaker() *gobreaker.CircuitBreaker {
	settings := gobreaker.Settings{
		Name:    "InvestmentAPI",
		Timeout: 5 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio >= 0.2
		},
	}
	return gobreaker.NewCircuitBreaker(settings)
}
