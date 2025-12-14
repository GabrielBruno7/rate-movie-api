package usecase

import "crud/domain/movie"

type MovieUsecase struct {
	repository movie.Repository
}

func NewMovieUsecase(repository movie.Repository) *MovieUsecase {
	return &MovieUsecase{repository: repository}
}

func (movieUsecase *MovieUsecase) ListPopularMovies(text string) ([]movie.Movie, error) {
	return movieUsecase.repository.ListPopularMovies(text)
}
