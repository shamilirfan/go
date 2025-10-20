package repo

import "fmt"

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email string, password string) (*User, error)
}

type userRepo struct {
	users []User
}

func NewUserRepo() UserRepo {
	return &userRepo{
		// users: []User{
		// 	{
		// 		ID: 1, FirstName: "Anisur",
		// 		LastName:    "Rahman",
		// 		Email:       "anisur@gmail.com",
		// 		Password:    "1234",
		// 		IsShopOwner: true,
		// 	},
		// },
	}
}

func (r *userRepo) Create(user User) (*User, error) {
	if user.ID != 0 {
		return &user, nil
	}

	user.ID = len(r.users) + 1
	r.users = append(r.users, user)
	return &user, nil
}

func (r *userRepo) Find(email string, password string) (*User, error) {
	for _, value := range r.users {
		if value.Email == email && value.Password == password {
			return &value, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}
