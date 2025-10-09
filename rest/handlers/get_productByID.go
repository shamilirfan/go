package handlers

import (
	"ecommerce/database"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productID")
	id, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please give me a valid id", 400)
		return
	}

	for index, value := range database.ProductList {
		if id == value.ID {
			json.NewEncoder(w).Encode(database.ProductList[index])
			return
		}
	}

	http.Error(w, "Not Found!!!", 404)
}
