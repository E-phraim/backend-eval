package handlers

import (
	"net/http"

	"github.com/e-phraim/backend-eval/utils"
	"github.com/gin-gonic/gin"
)

const csv = "./db/movies.csv"

func GetMovies(c *gin.Context){
	movies, err := utils.Reader(csv)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func GetMovieByID(c *gin.Context){
	
}