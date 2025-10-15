package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func(h *Handler)  DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	id, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please give me a valid id", 400)
		return
	}

	database.Delete(id)
	util.SendData(w, "Successfully deleted", 200)
}
