package usecase

import (
	"crud/domain/errs"
	"crud/domain/user"
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

func (authUsecase *AuthUsecase) Login(user *user.User) (string, error) {
	password := user.Password

	user, err := authUsecase.repository.LoadUserByEmail(user)
	if err != nil {
		return "", errs.NewWithCode(errs.ErrInvalidCredentials, err)
	}

	if user == nil {
		return "", errs.NewWithCode(errs.ErrInvalidCredentials, nil)
	}

	isPasswordValid := authUsecase.checkPassword(user.Password, password)
	if !isPasswordValid {
		return "", errs.NewWithCode(errs.ErrInvalidCredentials, nil)
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
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})

	return token.SignedString(secret)
}
