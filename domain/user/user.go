package user

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(id string, name string, email string, password string) *User {
	return &User{
		Id:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}
}
