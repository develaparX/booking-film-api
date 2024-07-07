package dto

type CreateGenreToMovieRequest struct {
	ID      string `json:"id"`
	GenreID string `json:"genre_id"`
	MovieID string `json:"movie_id"`
}

type GenreToMovieResponse struct {
	ID      string `json:"id"`
	GenreID string `json:"genre_id"`
	MovieID string `json:"movie_id"`
}
