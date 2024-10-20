package handler

import (
	"encoding/json"
	"fmt"
	"go-circuitbreaker/request"
	"net/http"
)

type Handler struct {
	investmentRequest *request.InvestmentRequest
}

func NewHandler(investmentRequest *request.InvestmentRequest) *Handler {
	return &Handler{
		investmentRequest: investmentRequest,
	}
}

func (h *Handler) GetInvestmentData(w http.ResponseWriter, r *http.Request) {
	data, err := h.investmentRequest.GetInvestmentData()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"status": "success",
		"data":   string(data),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
