package repo

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
	Find(email, password string) (*User, error)
	// Update(product User) (*User, error)
	// Delete(userID int) error
	// List() ([]*User, error)
}

type userRepo struct {
	users []User
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (r userRepo) Create(user User) (*User, error) {
	if user.ID != 0 {
		return &user, nil
	}

	user.ID = len(r.users) + 1
	r.users = append(r.users, user)
	return &user, nil
}

func (r userRepo) Find(email, password string) (*User, error) {
	for _, value := range r.users {
		if value.Email == email && value.Password == password {
			return &value, nil
		}
	}
	return nil, nil
}
