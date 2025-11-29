package user

import (
	"crud/http/handlers"
	"crud/infrastructure/database"
	"crud/usecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, db *sql.DB) {
	userUsecase := usecase.NewUserUsecase(database.NewUserDb(db))
	userHandler := handlers.NewUserHandler(userUsecase)

	router.POST("/user", userHandler.Create)
}
