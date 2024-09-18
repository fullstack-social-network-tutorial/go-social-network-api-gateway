package handlers

import (
	"net/http"
)

type AuthHandler struct {
	host          string
	apiGatewayKey string
}

func NewAuthHandler(host string, apiGatewayKey string) AuthHandler {
	return AuthHandler{host: host}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r http.Request) {
	//  Create a new HTTP Request
	request, err := http.NewRequest(http.MethodPost, h.host, r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Add Headers
	request.Header = r.Header.Clone()
	request.Header.Add("API-GATEWAY-KEY", h.apiGatewayKey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = response.Write(w)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
