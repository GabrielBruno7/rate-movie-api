package usecase

import (
	"crud/domain/user"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	repository user.Repository
}

func NewAuthUsecase(repository user.Repository) *AuthUsecase {
	return &AuthUsecase{
		repository: repository,
	}
}

func (authUsecase *AuthUsecase) Login(email string, password string) (string, error) {
	tempUser := &user.User{
		Email:    email,
		Password: password,
	}

	user, err := authUsecase.repository.LoadUserByEmail(tempUser)
	if err != nil {
		return "", errors.New("Credenciais inválidas")
	}

	if user == nil {
		return "", errors.New("Credenciais inválidas")
	}

	isPasswordValid := authUsecase.checkPassword(user.Password, password)
	if !isPasswordValid {
		return "", errors.New("Credenciais inválidas")
	}

	token, err := authUsecase.generateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (authUsecase *AuthUsecase) checkPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (authUsecase *AuthUsecase) generateToken(user *user.User) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Minute * 1).Unix(),
	})

	return token.SignedString(secret)
}
