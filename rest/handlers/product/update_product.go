package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	var newProduct repo.Product
	id, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please give me a valid id", http.StatusBadRequest)
		return
	}

	// decode
	decoder := json.NewDecoder(r.Body).Decode(&newProduct)

	if decoder != nil {
		http.Error(w, "Please give me a valid json.", http.StatusBadRequest)
		return
	}

	newProduct.ID = id

	h.productRepo.Update(newProduct)

	util.SendData(w, "Succesfully updated", http.StatusOK)
}
