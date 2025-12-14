CREATE TABLE IF NOT EXISTS rates (
    id UUID PRIMARY KEY NOT NULL,
    movie_name VARCHAR(255) NOT NULL,
    movie_tmdb_id VARCHAR(50) NOT NULL,
    movie_rate INTEGER NOT NULL CHECK (movie_rate >= 1 AND movie_rate <= 10),
    user_id UUID NOT NULL,
    comment TEXT,
    movie_image_path VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);


CREATE INDEX IF NOT EXISTS idx_rates_user_id ON rates(user_id);
CREATE INDEX IF NOT EXISTS idx_rates_movie_tmdb_id ON rates(movie_tmdb_id);
CREATE INDEX IF NOT EXISTS idx_rates_movie_rate ON rates(movie_rate);
