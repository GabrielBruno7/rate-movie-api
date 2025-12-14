package handlers

import (
	"crud/domain/errs"
	"crud/domain/rate"
	"crud/domain/user"
	"crud/http/dto"
	"crud/http/response"
	"crud/usecase"

	"github.com/gin-gonic/gin"
)

type RateHandler struct {
	usecase *usecase.RateUsecase
}

func NewRateHandler(usecase *usecase.RateUsecase) *RateHandler {
	return &RateHandler{usecase: usecase}
}

func (rateHandler *RateHandler) ActionRateMovie(context *gin.Context) {
	var request dto.RateMovieRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		response.BadRequest(context, errs.ErrInvalidBody)
		return
	}

	user := user.User{
		Email: context.GetString("email"),
	}

	rate := &rate.Rate{
		User:      user,
		Rate:      request.Rate,
		Name:      request.Name,
		TmdbId:    request.TmdbId,
		Comment:   request.Comment,
		ImagePath: request.ImagePath,
	}

	err := rateHandler.usecase.RateMovie(rate)
	if err != nil {
		response.HandleError(context, err)
		return
	}

	context.JSON(202, gin.H{"id": rate.ID})
}
