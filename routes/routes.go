package routes

import (
	"github.com/e-phraim/backend-eval/handlers"
	"github.com/gin-gonic/gin"
)

func MovieRoutes(engine *gin.Engine) {

	v1 := engine.Group("/v1")
	{
		v1.GET("/movies", handlers.GetMovies)
	}

}
