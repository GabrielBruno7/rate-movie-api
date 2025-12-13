package movie

import (
	"crud/http/handlers"
	"crud/http/middleware"
	"crud/infrastructure/database"
	"crud/usecase"
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
)

func RegisterMovieRoutes(router *gin.Engine, db *sql.DB) {
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	movieRepo := database.NewMovieDb(os.Getenv("TMDB_API_KEY"))
	movieUsecase := usecase.NewMovieUsecase(movieRepo)
	movieHandler := handlers.NewMovieHandler(movieUsecase)

	protected.GET("/movies/search", movieHandler.ActionListPopularMovies)
}
