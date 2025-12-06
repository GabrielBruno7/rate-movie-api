package movie

type Repository interface {
	ListPopularMovies() ([]Movie, error)
}
