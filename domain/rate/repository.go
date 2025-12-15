package rate

type Repository interface {
	RateMovie(rate *Rate) error
	UpdateRate(rate *Rate) error
	FindRateByTmdbId(rate *Rate) (*Rate, error)
	FindAllRatesByUser(rate *Rate) ([]*Rate, error)
	FindRateById(rate *Rate) (*Rate, error)
}
