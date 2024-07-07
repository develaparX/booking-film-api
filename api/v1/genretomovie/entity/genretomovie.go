package entity

type GenreToMovie struct {
	ID      string `json:"id"`
	GenreID string `json:"genre_id"`
	MovieID string `json:"movie_id"`
}
