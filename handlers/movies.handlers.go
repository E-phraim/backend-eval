package handlers

import (
	"fmt"
	"net/http"

	"github.com/e-phraim/backend-eval/db"
	"github.com/e-phraim/backend-eval/models"
	"github.com/e-phraim/backend-eval/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetMovies(c *gin.Context) {
	movies, err := utils.MoviesReader(db.Movie_csv)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func GetMovieByID(c *gin.Context) {
	id := c.Param("id")
	movieID, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("from parsing the id ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error: Invalid movie ID": err.Error()})
		return
	}

	// Read data from CSV files
	movies, err := utils.MoviesReader(db.Movie_csv)
	if err != nil {
		fmt.Println("from movies reader ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error: unable to read from movies reader ": err.Error()})
		return
	}

	// criticReviews, err := utils.CriticsReviewsReader(db.CriticsReviews_csv)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error: unable to read from critics reader": err.Error()})
	// 	return
	// }

	// userReviews, err := utils.UserReviewsReader(db.UsersReviews_csv)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// Find movie by ID
	var movie models.Movie
	found := false
	for _, m := range movies {
		if m.MovieID == movieID {
			movie = m
			found = true
			break
		}
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	// Find related reviews
	// var movieCriticReviews []models.CriticReview
	// for _, review := range criticReviews {
	// 	if review.MovieID == movieID {
	// 		movieCriticReviews = append(movieCriticReviews, review)
	// 	}
	// }

	// var movieUserReviews []models.UserReview
	// for _, review := range userReviews {
	// 	if review.MovieID == movieID {
	// 		movieUserReviews = append(movieUserReviews, review)
	// 	}
	// }

	// Combine data
	combinedResponse := struct {
		Movie         models.Movie
		CriticReviews []models.CriticReview
		UserReviews   []models.UserReview
	}{
		Movie: movie,
		// CriticReviews: movieCriticReviews,
		// UserReviews:   movieUserReviews,
	}

	c.JSON(http.StatusOK, combinedResponse)
}
