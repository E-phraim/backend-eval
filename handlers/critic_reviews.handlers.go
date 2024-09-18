package handlers

import (
	"net/http"

	"github.com/e-phraim/backend-eval/db"
	"github.com/e-phraim/backend-eval/utils"
	"github.com/gin-gonic/gin"
)

func GetReviewsFromCritics(c *gin.Context) {
	reviews, err := utils.CriticsReviewsReader(db.CriticsReviews_csv)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reviews)
}
