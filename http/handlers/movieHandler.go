package handlers

import (
	"crud/http/response"
	"crud/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	usecase *usecase.MovieUsecase
}

func NewMovieHandler(usecase *usecase.MovieUsecase) *MovieHandler {
	return &MovieHandler{usecase: usecase}
}

func (movieHandler *MovieHandler) ActionListPopularMovies(context *gin.Context) {
	movies, err := movieHandler.usecase.ListPopularMovies()
	if err != nil {
		response.HandleError(context, err)
		return
	}

	context.JSON(http.StatusOK, movies)
}
