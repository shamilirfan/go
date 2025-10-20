package product

import (
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productRepo.List()

	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusBadRequest)
		return
	}

	util.SendData(w, products, http.StatusOK)
}
