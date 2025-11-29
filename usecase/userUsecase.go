package usecase

import (
	"crud/domain/user"

	"github.com/google/uuid"
)

type UserUsecase struct {
	repository user.Repository
}

func NewUserUsecase(repository user.Repository) *UserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}

func (userUsecase *UserUsecase) CreateUser(name string, email string, password string) (string, error) {
	user := user.NewUser(
		uuid.New().String(),
		name,
		email,
		password,
	)

	err := userUsecase.repository.Create(user)

	return user.Id, err
}
