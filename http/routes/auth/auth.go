package auth

import (
	"crud/http/handlers"
	"crud/infrastructure/database"
	"crud/usecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine, db *sql.DB) {
	authUsecase := usecase.NewAuthUsecase(database.NewUserDb(db))
	authHandler := handlers.NewAuthHandler(authUsecase)

	router.POST("/login", authHandler.Login)
}
