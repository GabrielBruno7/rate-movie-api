package movie

type Repository interface {
	ListPopularMovies(text string) ([]Movie, error)
}
