package routes

import (
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func initMovieRouter(router *gin.Engine) {
	movieHandler := handlers.NewMovieHandler()

	movieRouter := router.Group("/movies")
	{
		movieRouter.GET("", movieHandler.AllMovies)
		movieRouter.GET("/:id", movieHandler.MovieDetail)
		movieRouter.GET("/filter", movieHandler.FilterMovies)
		movieRouter.GET("/upcoming", movieHandler.UpcomingMovies)
	}
}