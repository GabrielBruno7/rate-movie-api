package dto

type RateMovieRequest struct {
	Rate      int    `json:"nota" binding:"required,min=1,max=10"`
	Comment   string `json:"comentario" binding:"required,max=100"`
	Name      string `json:"nome" binding:"required"`
	TmdbId    string `json:"tmdb_id" binding:"required"`
	ImagePath string `json:"caminho_imagem" binding:"required"`
}
