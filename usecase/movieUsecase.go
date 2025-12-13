package usecase

import "crud/domain/movie"

type MovieUsecase struct {
	repository movie.Repository
}

func NewMovieUsecase(r movie.Repository) *MovieUsecase {
	return &MovieUsecase{repository: r}
}

func (movieUsecase *MovieUsecase) ListPopularMovies(text string) ([]movie.Movie, error) {
	return movieUsecase.repository.ListPopularMovies(text)
}
