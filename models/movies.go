package models

import (
	"github.com/google/uuid"
)

type AddMovie struct {
	MovieTitle    string `json:"movie_title"`
	Year          string `json:"year"`
	URL           string `json:"url"`
	CriticScore   string `json:"critic_score"`
	AudienceScore string `json:"audience_score"`
}

type Movie struct {
	MovieID       uuid.UUID `json:"movie_id"`
	MovieTitle    string    `json:"movie_title"`
	Year          int       `json:"year"`
	URL           string    `json:"url"`
	CriticScore   string    `json:"critic_score"`
	AudienceScore string    `json:"audience_score"`
}
