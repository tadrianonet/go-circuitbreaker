package main

import (
	"fmt"
	"go-circuitbreaker/config"
	"go-circuitbreaker/handler"
	"go-circuitbreaker/request"
	"net/http"
)

func main() {

	circuitBreaker := config.NewCircuitBreaker()
	investmentRequest := request.NewInvestmentRequest(circuitBreaker)

	investmentHandler := handler.NewHandler(investmentRequest)

	http.HandleFunc("/api/v1/investment", investmentHandler.GetInvestmentData)
	fmt.Println("Investment API running on port 8080")
	http.ListenAndServe(":8080", nil)
}
