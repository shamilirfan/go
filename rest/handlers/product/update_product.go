package product


import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

func(h *Handler)  UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	var newProduct database.Product
	id, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please give me a valid id", 400)
		return
	}

	// decode
	decoder := json.NewDecoder(r.Body).Decode(&newProduct)

	if decoder != nil {
		http.Error(w, "Please give me a valid json.", 400)
		return
	}

	newProduct.ID = id

	database.Update(newProduct)

	util.SendData(w, "Succesfully updated", 200)
}
