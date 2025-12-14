package handlers

import (
	"crud/http/dto"
	"crud/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(usecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (h *UserHandler) Create(context *gin.Context) {
	var request dto.CreateUserRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(400, gin.H{"error": "Dados inválidos"})
		return
	}

	id, err := h.usecase.CreateUser(
		request.Name,
		request.Email,
		request.Password,
	)
	if err != nil {
		context.JSON(500, gin.H{
			"error":      "Erro ao criar usuário", //TODO: Alterar para utilizar erro novo
			"stacktrace": err.Error(),
		})
		return
	}

	context.JSON(201, gin.H{"id": id})
}
