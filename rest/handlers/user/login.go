package user

import (
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin

	err := json.NewDecoder(r.Body).Decode(&reqLogin)

	if err != nil {
		http.Error(w, "Please give me a valid json", http.StatusBadRequest)
		return
	}

	user := database.Find(reqLogin.Email, reqLogin.Password)

	if user == nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	cnf := config.GetConfig()
	accessToken, err := util.CreateJwt(cnf.JwtSecretKey, util.Payload{
		Sub:         user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopOwner: user.IsShopOwner,
	})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, accessToken, http.StatusCreated)
}
