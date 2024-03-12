package controllers

import (
	"customermod/models"
	"encoding/json"
	"net/http"
)

func BuildResponse(rw http.ResponseWriter, data interface{}, err error) {
    response := models.Response{}
	if err != nil {
		response.Error = err.Error()
	} else {
		response.Load = data
	}

	jsonResponse, _ := json.Marshal(response)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonResponse)
}