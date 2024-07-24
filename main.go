package main

import (
	"github.com/e-phraim/backend-eval/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/movies", handlers.GetMovies)

	engine.Run(":7000")
}
