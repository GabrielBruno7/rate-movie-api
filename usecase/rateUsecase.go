package usecase

import (
	"crud/domain/errs"
	"crud/domain/rate"

	"github.com/google/uuid"
)

type RateUsecase struct {
	repository  rate.Repository
	UserUsecase *UserUsecase
}

func NewRateUsecase(repository rate.Repository, userUsecase *UserUsecase) *RateUsecase {
	return &RateUsecase{
		repository:  repository,
		UserUsecase: userUsecase,
	}
}

func (rateUsecase *RateUsecase) RateMovie(rate *rate.Rate) error {
	_, err := rateUsecase.loadUser(rate)
	if err != nil {
		return err
	}

	existingRate, err := rateUsecase.repository.FindRateByTmdbId(rate)
	if err != nil {
		return err
	}

	if existingRate != nil {
		return rateUsecase.repository.UpdateRate(existingRate)
	}

	rate.ID = uuid.New().String()

	return rateUsecase.repository.RateMovie(rate)
}

func (rateUsecase *RateUsecase) ListRates(rate *rate.Rate) ([]*rate.Rate, error) {
	_, err := rateUsecase.loadUser(rate)
	if err != nil {
		return nil, err
	}

	return rateUsecase.repository.FindAllRatesByUser(rate)
}

func (rateUsecase *RateUsecase) loadUser(rate *rate.Rate) (*rate.Rate, error) {
	user, err := rateUsecase.UserUsecase.LoadUserByEmail(&rate.User)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errs.NewWithCode(errs.ErrUserNotFound, nil)
	}

	rate.User = *user

	return rate, nil
}
