package utils

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/e-phraim/backend-eval/db"
	"github.com/e-phraim/backend-eval/models"
	"github.com/google/uuid"
)

func MoviesReader(filepath string) ([]models.Movie, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var movies []models.Movie
	for _, record := range records[1:] {
		year, _ := strconv.Atoi(record[2])
		movies = append(movies, models.Movie{
			MovieID:       uuid.MustParse(record[0]),
			MovieTitle:    record[1],
			Year:          year,
			URL:           record[3],
			CriticScore:   record[4],
			AudienceScore: record[5],
		})
	}

	return movies, nil
}

func ParseBoolValue(value string) (bool, error) {
	value = strings.ToLower(value)

	if value == "" {
		return false, nil
	}

	return strconv.ParseBool(value)
}

func CriticsReviewsReader(filepath string) ([]models.CriticReview, error) {
	file, err := os.Open(db.CriticsReviews_csv)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var criticsReview []models.CriticReview
	for _, record := range records[1:] {

		reviewID, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Error parsing reviewId: %v", err)
			return nil, err
		}

		isFresh, err := ParseBoolValue(record[6])
		if err != nil {
			log.Printf("Error parsing isFresh: %v", err)
			continue
		}

		// isRotten, err := strconv.ParseBool(record[8])
		// if err != nil {
		// 	log.Printf("Error parsing isRotten: %v", err)
		// 	return nil, err
		// }

		// isRTURL, err := strconv.ParseBool(record[9])
		// if err != nil {
		// 	log.Printf("Error parsing isRTURL: %v", err)
		// 	return nil, err
		// }

		// isTopCritic, err := strconv.ParseBool(record[9])
		// if err != nil {
		// 	log.Printf("Error parsing isTopCritic: %v", err)
		// 	return nil, err
		// }

		criticsReview = append(criticsReview, models.CriticReview{
			// ReviewID:        uuid.MustParse(record[0]),
			ReviewID:      reviewID,
			MovieID:       uuid.MustParse(record[1]),
			CreationDate:  record[3],
			CriticName:    record[4],
			CriticPageURL: record[5],
			ReviewState:   record[6],
			IsFresh:       isFresh,
			// IsRotten:        isRotten,
			// IsRTURL:         isRTURL,
			// IsTopCritic:     isTopCritic,
			PublicationURL:  record[10],
			PublicationName: record[11],
			ReviewURL:       record[12],
			Quote:           record[13],
			ScoreSentiment:  record[14],
			OriginalScore:   record[15],
		})
	}

	return criticsReview, nil
}

func UserReviewsReader(filepath string) ([]models.UserReview, error) {
	file, err := os.Open(db.UsersReviews_csv)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var reviews []models.UserReview
	for _, record := range records[1:] {
		// Parse the UUID for MovieID
		movieID, err := uuid.Parse(record[0])
		if err != nil {
			return nil, err
		}

		// Parse the rating
		rating, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, err
		}

		// Parse the ReviewID (index 3)
		reviewID, err := strconv.Atoi(record[3])
		if err != nil {
			log.Printf("Error parsing reviewId: %v", err)
			return nil, err
		}

		// Parse the boolean values
		isVerified, err := strconv.ParseBool(record[4])
		if err != nil {
			return nil, err
		}
		isSuperReviewer, err := strconv.ParseBool(record[5])
		if err != nil {
			return nil, err
		}
		hasSpoilers, err := strconv.ParseBool(record[6])
		if err != nil {
			return nil, err
		}
		hasProfanity, err := strconv.ParseBool(record[7])
		if err != nil {
			return nil, err
		}

		// Parse the score
		score, err := strconv.ParseFloat(record[8], 64)
		if err != nil {
			return nil, err
		}

		// Create the UserReview struct and append it to the reviews slice
		review := models.UserReview{
			MovieID:         movieID,
			Rating:          rating,
			Quote:           record[2], // Quote at index 2
			ReviewID:        reviewID,  // Correct ReviewID index (3)
			IsVerified:      isVerified,
			IsSuperReviewer: isSuperReviewer,
			HasSpoilers:     hasSpoilers,
			HasProfanity:    hasProfanity,
			Score:           score,
			CreationDate:    record[9],  // Creation date at index 9
			UserDisplayName: record[10], // User display name at index 10
			UserRealm:       record[11], // User realm at index 11
			UserID:          record[12], // User ID at index 12
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

func Writer(filepath string, movies []models.Movie) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Movie ID", "Movie Title", "Year", "URL", "Critic Score", "Audience Score"})

	for _, movie := range movies {
		writer.Write([]string{movie.MovieID.String(), movie.MovieTitle, strconv.Itoa(movie.Year), movie.URL, movie.CriticScore, movie.AudienceScore})
	}

	return nil
}
