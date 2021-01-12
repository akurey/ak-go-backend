package controllers

import (
	"api-test/db"
	"api-test/entities"
	"api-test/models"
	"encoding/json"
	fStructs "github.com/fatih/structs"
	"net/http"
	"time"
)

func ListProducts(response http.ResponseWriter, request *http.Request) {
	type error interface {
		Error() string
	}
	var (
		productsStruct             []models.Products
		responseMultiDataStruct entities.ResponseMultiDataStruct
	)
	if err := db.DB.Find(&productsStruct).Error; err != nil {
		response.WriteHeader(http.StatusBadRequest)
		responseMultiDataStruct.Status = false
		responseMultiDataStruct.Message = "Error listing products"
		responseMultiDataStruct.Result = nil
	} else {
		var results = make([]map[string]interface{}, len(productsStruct))
		for i := 0; i < len(productsStruct); i++ {
			results[i] = fStructs.Map(productsStruct[i])
		}
		responseMultiDataStruct.Status = true
		responseMultiDataStruct.Message = "Success"
		responseMultiDataStruct.Result = results
	}
	json.NewEncoder(response).Encode(&responseMultiDataStruct)
}

func CreateProduct(response http.ResponseWriter, request *http.Request) {
	type error interface {
		Error() string
	}
	var (
		productStruct    models.Products
		responseStruct entities.ResponseStruct
	)
	response.Header().Set("Content-Type", "application/json")
	currentDate := time.Now()
	if err := json.NewDecoder(request.Body).Decode(&productStruct); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		responseStruct.Status = false
		responseStruct.Message = err.Error()
		responseStruct.Result = nil
		json.NewEncoder(response).Encode(&responseStruct)
		return
	}
	productStruct.CreatedAt = currentDate
	productStruct.UpdatedAt = currentDate
	if err := db.DB.Create(&productStruct).Error; err != nil {
		response.WriteHeader(http.StatusBadRequest)
		responseStruct.Status = false
		responseStruct.Message = "Error creating product"
		responseStruct.Result = nil
	} else {
		responseStruct.Status = true
		responseStruct.Message = "Success"
		responseStruct.Result = fStructs.Map(&productStruct)
	}
	json.NewEncoder(response).Encode(&responseStruct)
}