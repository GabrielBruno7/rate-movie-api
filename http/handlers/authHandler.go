package handlers

import (
	"crud/domain/errs"
	"crud/domain/user"
	"crud/http/dto"
	"crud/http/response"
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
		response.BadRequest(context, errs.ErrInvalidBody)
		return
	}

	user := &user.User{
		Email:    request.Email,
		Password: request.Password,
	}

	token, err := authHandler.usecase.Login(user)
	if err != nil {
		response.HandleError(context, err)
		return
	}

	context.JSON(200, gin.H{"token": token})
}
