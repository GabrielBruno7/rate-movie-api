package tmdb

import (
	"crud/domain/errs"
	"fmt"
	"os"
	"strconv"

	"github.com/go-resty/resty/v2"
)

type TMDBClient struct {
	client  *resty.Client
	apiKey  string
	baseURL string
}

func NewTMDBClient(apiKey string) *TMDBClient {
	return &TMDBClient{
		client:  resty.New(),
		apiKey:  apiKey,
		baseURL: os.Getenv("TMDB_BASE_URL"),
	}
}

func (tmdbClient *TMDBClient) FetchPopularMovies(page int) (*PopularMoviesResponse, error) {
	var result PopularMoviesResponse

	response, err := tmdbClient.client.R().
		SetQueryParams(map[string]string{
			"api_key":  tmdbClient.apiKey,
			"language": "pt-BR",
			"page":     strconv.Itoa(page),
		}).
		SetResult(&result).
		Get(tmdbClient.baseURL + "/movie/popular")

	if err != nil {
		return nil, errs.NewWithCode(errs.ErrTMDBConnection, err)
	}

	if response.IsError() {
		return nil, errs.NewWithCode(errs.ErrTMDBAPIError, fmt.Errorf(
			"HTTP %d: %s", response.StatusCode(), response.Status(),
		))
	}

	return &result, nil
}
