package user

import (
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("userRepo:", h.userRepo)
	var req ReqLogin

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.SendError(w, "Please give me a valid json", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.Find(req.Email, req.Password)
	if err != nil {
		util.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	accessToken, err := util.CreateJwt(h.config.JwtSecretKey, util.Payload{
		Sub:       user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})

	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, accessToken, http.StatusCreated)
}
