package pkg

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int
	Payload interface{}
	Success bool
}

func HttpError(w http.ResponseWriter, code int, err string) {
	errorPayload := Response{
		Code:    code,
		Payload: err,
		Success: false,
	}
	json.NewEncoder(w).Encode(errorPayload)
}

func HttpSuccess(w http.ResponseWriter, code int, payload interface{}) {
	successPayload := Response{
		Code:    code,
		Payload: payload,
		Success: true,
	}

	json.NewEncoder(w).Encode(successPayload)
}
