package database

import (
	"crud/domain/rate"
	"database/sql"
)

type RateDb struct {
	db *sql.DB
}

func NewRateDb(db *sql.DB) *RateDb {
	return &RateDb{db: db}
}

func (persistence *RateDb) RateMovie(rate *rate.Rate) error {
	query := `
	    INSERT INTO rates (
			id,
			movie_rate,
			movie_name,
			movie_tmdb_id,
			user_id,
			comment,
			movie_image_path
	    ) VALUES ($1, $2, $3, $4, $5, $6, $7)
	    RETURNING id
	`

	_, err := persistence.db.Exec(query,
		rate.ID,
		rate.Rate,
		rate.Name,
		rate.TmdbId,
		rate.User.Id,
		rate.Comment,
		rate.ImagePath,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repository *RateDb) FindRateByTmdbId(rate *rate.Rate) (*rate.Rate, error) {
	query := `
		SELECT id
		FROM rates
		WHERE movie_tmdb_id = $1 AND user_id = $2
		LIMIT 1
	`

	err := repository.db.QueryRow(query, rate.TmdbId, rate.User.Id).Scan(
		&rate.ID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return rate, nil
}
func (repository *RateDb) UpdateRate(rate *rate.Rate) error {
	query := `
		UPDATE rates
		SET 
			movie_rate = $1,
			comment = $2,
			movie_name = $3,
			movie_image_path = $4,
			updated_at = NOW()
		WHERE 
			id = $5 AND user_id = $6
	`

	_, err := repository.db.Exec(query,
		rate.Rate,
		rate.Comment,
		rate.Name,
		rate.ImagePath,
		rate.ID,
		rate.User.Id,
	)

	if err != nil {
		return err
	}

	return nil
}
