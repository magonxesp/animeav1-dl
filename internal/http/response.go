package http

import (
	"encoding/json"
	"net/http"

	"github.com/magonxesp/animeav1-dl/internal/common"
)

type ErrorResponse struct {
	message string
}

func RespondJSON(writter http.ResponseWriter, status int, payload any) {
	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(status)

	if err := json.NewEncoder(writter).Encode(payload); err != nil {
		common.Logger.Error("error al codificar la respuesta JSON", "error", err)
	}
}

func RespondJSONError(writter http.ResponseWriter, status int, message string) {
	body := ErrorResponse{message}
	RespondJSON(writter, status, body)
}
