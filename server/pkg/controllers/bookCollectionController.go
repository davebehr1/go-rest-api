package controllers

import (
	"encoding/json"
	"net/http"
)

func Test(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("api is alive")
}
