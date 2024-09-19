package routes

import (
	"github.com/e-phraim/backend-eval/handlers"
	"github.com/gin-gonic/gin"
)

func MovieRoutes(engine *gin.Engine) {

	v1 := engine.Group("/v1")
	{
		// movies
		v1.GET("/movies", handlers.GetMovies)
		v1.GET("/movie/:id", handlers.GetMovieByID)

		// critic reviews
		v1.GET("/reviews", handlers.GetReviewsFromCritics)
		// v1.GET("/review/:id", handlers.GetReviewByID)

		// user reviews[comments]
		v1.GET("/comments", handlers.GetComments)
		v1.GET("/comment/:id", handlers.GetCommentByID)
	}

}
