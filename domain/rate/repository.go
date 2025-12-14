package rate

type Repository interface {
	RateMovie(rate *Rate) error
	UpdateRate(rate *Rate) error
	FindRateByTmdbId(rate *Rate) (*Rate, error)
}
