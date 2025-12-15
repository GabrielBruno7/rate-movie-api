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

func (rateHandler *RateHandler) ActionListRates(context *gin.Context) {
	user := user.User{
		Email: context.GetString("email"),
	}

	rate := &rate.Rate{
		User: user,
	}

	rates, err := rateHandler.usecase.ListRates(rate)
	if err != nil {
		response.HandleError(context, err)
		return
	}

	mappedRates := make([]map[string]interface{}, 0, len(rates))
	for _, r := range rates {
		rateItem := map[string]interface{}{
			"ID":        r.ID,
			"Name":      r.Name,
			"TmdbId":    r.TmdbId,
			"Rate":      r.Rate,
			"Comment":   r.Comment,
			"ImagePath": r.ImagePath,
		}

		mappedRates = append(mappedRates, rateItem)
	}

	context.JSON(200, mappedRates)
}

func (rateHandler *RateHandler) ActionRateDetails(context *gin.Context) {
	user := user.User{
		Email: context.GetString("email"),
	}

	rate := &rate.Rate{
		ID:   context.Param("id"),
		User: user,
	}

	rate, err := rateHandler.usecase.LoadRateById(rate)
	if err != nil {
		response.BadRequest(context, errs.ErrRateNotFound)
		return
	}

	rateItem := map[string]interface{}{
		"ID":        rate.ID,
		"Name":      rate.Name,
		"TmdbId":    rate.TmdbId,
		"Rate":      rate.Rate,
		"Comment":   rate.Comment,
		"ImagePath": rate.ImagePath,
	}

	context.JSON(200, rateItem)
}
