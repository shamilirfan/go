package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// new variable declare
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
