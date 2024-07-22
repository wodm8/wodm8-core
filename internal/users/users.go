package users

type Users struct {
	ID        string `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
}

func NewUsers(id, firstname, lastname, email, password string) (Users, error) {
	return Users{
		ID:        id,
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		Password:  password,
	}, nil
}

type UsersRepository interface {
	Save(user Users) error
	First(email string) (*Users, error)
}
