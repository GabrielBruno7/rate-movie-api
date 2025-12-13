package handlers

import (
	"crud/domain/errs"
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
	text := context.Query("filme")
	if text == "" {
		response.BadRequest(context, errs.ErrMissingParameter)
		return
	}

	movies, err := movieHandler.usecase.ListPopularMovies(text)
	if err != nil {
		response.HandleError(context, err)
		return
	}

	context.JSON(http.StatusOK, movies)
}
