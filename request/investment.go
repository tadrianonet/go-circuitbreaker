package request

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sony/gobreaker"
)

type InvestmentRequest struct {
	cb     *gobreaker.CircuitBreaker
	client http.Client
	url    string
}

func NewInvestmentRequest(cb *gobreaker.CircuitBreaker) *InvestmentRequest {
	return &InvestmentRequest{
		cb:     cb,
		client: http.Client{},
		url:    "https://go-invest-example.onrender.com/api/v1/investment",
	}
}

func (r *InvestmentRequest) GetInvestmentData() ([]byte, error) {
	switch r.cb.State() {
	case gobreaker.StateClosed:
		fmt.Println("Circuit Breaker Status: FECHADO")
	case gobreaker.StateOpen:
		fmt.Println("Circuit Breaker Status: ABERTO")
	case gobreaker.StateHalfOpen:
		fmt.Println("Circuit Breaker Status: SEMI-ABERTO")
	}

	body, err := r.cb.Execute(func() (interface{}, error) {

		req, err := http.NewRequest(http.MethodGet, r.url, nil)
		if err != nil {
			return nil, err
		}

		resp, err := r.client.Do(req)
		if err != nil {
			fmt.Println("Erro ao chamar a API de investimentos:", err)
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Erro na resposta da API de investimentos: %v\n", resp.Status)
			return nil, fmt.Errorf("erro na resposta: %v", resp.Status)
		}

		return io.ReadAll(resp.Body)
	})

	if err != nil {
		fmt.Printf("Erro ao buscar dados: %v\n", err)
		return nil, err
	}

	return body.([]byte), nil
}
