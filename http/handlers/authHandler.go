package handlers

import (
	"crud/http/dto"
	"crud/usecase"

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

	token, err := authHandler.usecase.Login(request.Email, request.Password)
	if err != nil {
		context.JSON(401, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"token": token})
}
