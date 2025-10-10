package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	id, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please give me a valid id", 400)
		return
	}

	product := database.Get(id)

	if product == nil {
		util.SendData(w, "Product Not Found!!!", 404)
		return
	}

	util.SendData(w, product, 200)
}
