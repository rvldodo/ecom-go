package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ResponseFormat struct {
	Status     bool        `json:"status"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

var Validate = validator.New()

func ParesJSON(r *http.Request, payload interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("Missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	statusRes := true
	if status >= 400 {
		statusRes = false
	}

	return json.NewEncoder(w).Encode(&ResponseFormat{
		Status:     statusRes,
		StatusCode: status,
		Data:       v,
	})
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
