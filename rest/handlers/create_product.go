package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product

	// Decode
	err := json.NewDecoder(r.Body).Decode(&newProduct)

	if err != nil {
		http.Error(w, "Please give me a valid json", 400)
		return
	}

	createdProduct := database.Store(newProduct)
	util.SendData(w, createdProduct, 201)
}
