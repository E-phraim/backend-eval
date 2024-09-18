package main

import (
	"github.com/e-phraim/backend-eval/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.MovieRoutes(r)

	r.Run("localhost:7100")
}
