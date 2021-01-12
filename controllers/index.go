package controllers

import (
	"api-test/entities"
	"encoding/json"
	"net/http"
)

// IndexHandlerGET handler get for index request
func IndexHandlerGET(w http.ResponseWriter, r *http.Request) {
	var (
		responseStruct entities.ResponseStruct
	)
	result := &responseStruct
	w.Header().Set("Content-Type", "application/json")
	result.Status = true
	result.Result = map[string]interface{}{
		"connection": true,
	}
	json.NewEncoder(w).Encode(result)
}
