package database

import (
	"crud/domain/movie"
	"crud/internal/tmdb"
	"os"
	"strconv"
)

type TMDbClient struct {
	APIKey  string
	BaseURL string
}

type TMDbMovieRepository struct {
	client *tmdb.TMDBClient
}

func NewTMDbMovieRepository(client *tmdb.TMDBClient) movie.Repository {
	return &TMDbMovieRepository{client: client}
}

func NewMovieDb(apiKey string) movie.Repository {
	client := tmdb.NewTMDBClient(apiKey)
	return NewTMDbMovieRepository(client)
}

func (tmdbMovieRepository *TMDbMovieRepository) ListPopularMovies() ([]movie.Movie, error) {
	response, err := tmdbMovieRepository.client.FetchPopularMovies(1)
	if err != nil {
		return nil, err // Repassa o erro específico do client (ErrTMDBConnection ou ErrTMDBAPIError)
	}

	movies := make([]movie.Movie, len(response.Results))
	for i, tmdbMovie := range response.Results {
		movies[i] = movie.Movie{
			ID:          strconv.Itoa(tmdbMovie.ID),
			Title:       tmdbMovie.Title,
			Overview:    tmdbMovie.Overview,
			PosterURL:   os.Getenv("TMDB_BASE_URL_POSTER") + tmdbMovie.PosterPath,
			ReleaseDate: tmdbMovie.ReleaseDate,
			Rating:      tmdbMovie.VoteAverage,
		}
	}

	return movies, nil
}
