package user

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, "Please give me a valid json", http.StatusBadRequest)
		return
	}

	createdUser := newUser.Store()
	util.SendData(w, createdUser, http.StatusCreated)
}
