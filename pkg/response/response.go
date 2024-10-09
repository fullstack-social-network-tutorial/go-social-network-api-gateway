package response

import (
	"encoding/json"
	"go-service/pkg/logger"
	"io"
	"net/http"
)

func Response(w http.ResponseWriter, statusCode int, body interface{}) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBody)
}

func ForwardResponse(w http.ResponseWriter, resp *http.Response, logger *logger.Logger) {
	defer resp.Body.Close()

	// Set the status code from the called request
	w.WriteHeader(resp.StatusCode)

	// Copy the headers from the called request's response to your response
	for name, values := range resp.Header {
		w.Header()[name] = values
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.LogError(err.Error(), nil)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write(body)
}
