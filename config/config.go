package config

import (
	"time"

	"github.com/sony/gobreaker"
)

func NewCircuitBreaker() *gobreaker.CircuitBreaker {
	settings := gobreaker.Settings{
		Name:    "InvestmentAPI",
		Timeout: 5 * time.Second, // Tempo antes de tentar abrir o circuito novamente
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// Abre o circuito após 3 falhas consecutivas
			return counts.TotalFailures >= 3
		},
	}
	return gobreaker.NewCircuitBreaker(settings)
}
