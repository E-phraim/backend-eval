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

type CriticReview struct {
	ReviewID        int       `json:"reviewId"`
	MovieID         uuid.UUID `json:"movieId"`
	CreationDate    string    `json:"creationDate"`
	CriticName      string    `json:"criticName"`
	CriticPageURL   string    `json:"criticPageUrl"`
	ReviewState     string    `json:"reviewState"`
	IsFresh         bool      `json:"isFresh"`
	IsRotten        bool      `json:"isRotten"`
	IsRTURL         bool      `json:"isRtUrl"`
	IsTopCritic     bool      `json:"isTopCritic"`
	PublicationURL  string    `json:"publicationUrl"`
	PublicationName string    `json:"publicationName"`
	ReviewURL       string    `json:"reviewUrl"`
	Quote           string    `json:"quote"`
	ScoreSentiment  string    `json:"scoreSentiment"`
	OriginalScore   string    `json:"originalScore"`
}

type UserReview struct {
	MovieID         uuid.UUID `json:"movieId"`
	Rating          float64   `json:"rating"`
	Quote           string    `json:"quote"`
	ReviewID        string    `json:"reviewId"`
	IsVerified      bool      `json:"isVerified"`
	IsSuperReviewer bool      `json:"isSuperReviewer"`
	HasSpoilers     bool      `json:"hasSpoilers"`
	HasProfanity    bool      `json:"hasProfanity"`
	Score           float64   `json:"score"`
	CreationDate    string    `json:"creationDate"`
	UserDisplayName string    `json:"userDisplayName"`
	UserRealm       string    `json:"userRealm"`
	UserID          string    `json:"userId"`
}

type Response struct {
	Movie         Movie          `json:"movie"`
	CriticReviews []CriticReview `json:"critic_review"`
	UserReviews   []UserReview   `json:"user_review"`
}
