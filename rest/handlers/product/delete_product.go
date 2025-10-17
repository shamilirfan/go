package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	id, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please give me a valid id", http.StatusBadRequest)
		return
	}

	h.productRepo.Delete(id)
	util.SendData(w, "Successfully deleted", http.StatusOK)
}
