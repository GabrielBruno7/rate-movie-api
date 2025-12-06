package handlers

import (
	"crud/http/dto"
	"crud/usecase"
	"crud/domain/user"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	usecase *usecase.AuthUsecase
}

func NewAuthHandler(usecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		usecase: usecase,
	}
}

func (authHandler *AuthHandler) Login(context *gin.Context) {
	var request dto.LoginRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(400, gin.H{"error": "Dados inválidos"})
		return
	}

	user := &user.User{
		Email:    request.Email,
		Password: request.Password,
	}

	token, err := authHandler.usecase.Login(user)
	if err != nil {
		context.JSON(401, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"token": token})
}
