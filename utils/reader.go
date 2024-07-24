package utils

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/e-phraim/backend-eval/models"
	"github.com/google/uuid"
)

func Reader(filepath string) ([]models.Movie, error) {
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
