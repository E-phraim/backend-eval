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

func GetComments(c *gin.Context) {
	comments, err := utils.UserReviewsReader(db.UsersReviews_csv)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading comments from csv" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func GetCommentByID(c *gin.Context) {
	id := c.Param("id")
	commentID, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("from parsing the id ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error: invalid comment ID": err.Error()})
		return
	}

	comments, err := utils.UserReviewsReader(db.UsersReviews_csv)
	if err != nil {
		fmt.Println("error from comments reader ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error: unable to read from comments reader ": err.Error()})
		return
	}

	var comment models.UserReview
	found := false
	for _, c := range comments {
		if c.MovieID == commentID {
			comment = c
			found = true
			break
		}
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "comment not found"})
		return
	}

	// Return the single comment
	c.JSON(http.StatusOK, comment)
}
