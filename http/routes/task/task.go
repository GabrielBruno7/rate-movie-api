package task

import (
	"crud/http/handlers"
	"crud/http/middleware"
	"crud/infrastructure/database"
	"crud/usecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.Engine, db *sql.DB) {
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	taskRepo := database.NewTaskDb(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskUsecase)

	protected.GET("/tasks", taskHandler.List)
	protected.POST("/tasks", taskHandler.Create)
}
