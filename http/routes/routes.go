package routes

import (
	"crud/http/routes/auth"
	"crud/http/routes/task"
	"crud/http/routes/user"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *sql.DB) *gin.Engine {
	router := gin.Default()

	task.RegisterTaskRoutes(router, db)
	user.RegisterUserRoutes(router, db)
	auth.RegisterAuthRoutes(router, db)

	return router
}
