package pkg

import (
	"encoding/json"
	"net/http"
)

type Success struct {
	Code    int
	Payload interface{}
	Success bool
}

type Error struct {
	Code         int
	ErrorMessage string
	Success      bool
}

func HttpError(w http.ResponseWriter, code int, err string) {
	errorPayload := Error{
		Code:         code,
		ErrorMessage: err,
		Success:      false,
	}
	json.NewEncoder(w).Encode(errorPayload)
}

func HttpSuccess(w http.ResponseWriter, code int, payload interface{}) {
	successPayload := Success{
		Code:    code,
		Payload: payload,
		Success: true,
	}

	json.NewEncoder(w).Encode(successPayload)
}
