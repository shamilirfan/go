package user

import (
	"ecommerce/config"
	"ecommerce/repo"
)

type Handler struct {
	config   config.Config
	userRepo repo.UserRepo
}

func NewHandler(config config.Config, userRepo repo.UserRepo) *Handler {
	return &Handler{
		config:   config,
		userRepo: userRepo,
	}
}
