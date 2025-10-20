package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// new variable declare
	var newProduct repo.Product

	// Decode
	err := json.NewDecoder(r.Body).Decode(&newProduct)

	if err != nil {
		util.SendError(w, "Please give me a valid json", http.StatusBadRequest)
		return
	}

	createdProduct, err := h.productRepo.Create(newProduct)

	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, createdProduct, http.StatusCreated)
}
