package handlers

import (
	"go-service/pkg/logger"
	"go-service/pkg/response"
	"net/http"
)

type AuthHandler struct {
	address       string
	apiGatewayKey string
	logger        *logger.Logger
}

func NewAuthHandler(address string, apiGatewayKey string, logger *logger.Logger) AuthHandler {
	return AuthHandler{address: address, apiGatewayKey: apiGatewayKey, logger: logger}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	//  Create a new HTTP Request
	request, err := http.NewRequest(http.MethodPost, h.address+"/login", r.Body)
	if err != nil {
		h.logger.LogError(err.Error(), nil)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Add Headers
	request.Header = r.Header.Clone()
	request.Header.Add("API-GATEWAY-KEY", h.apiGatewayKey)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		h.logger.LogError(err.Error(), nil)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response.ForwardResponse(w, resp, h.logger)
}
