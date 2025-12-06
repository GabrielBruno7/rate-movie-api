package auth

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuth(email string, password string) *Auth {
	return &Auth{
		Email:    email,
		Password: password,
	}
}
