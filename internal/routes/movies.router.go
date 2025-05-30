package routes

import (
	"main/internal/handlers"
	"main/internal/middlewares"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func initMovieRouter(router *gin.Engine, pg *pgxpool.Pool, rdc *redis.Client, mw *middlewares.Middleware) {
	movieRepo := repositories.NewMovieRepository(pg, rdc)
	movieHandler := handlers.NewMovieHandler(movieRepo)

	movieRouter := router.Group("/movies")
	movieRouter.GET("", movieHandler.AllMovies)
	movieRouter.GET("/:id", movieHandler.MovieDetail)
	movieRouter.GET("/filter", movieHandler.FilterMovies)
	movieRouter.GET("/now-playing", movieHandler.NowPlayingMovies)
	movieRouter.GET("/upcoming", movieHandler.UpcomingMovies)

	adminRouter := router.Group("/admin/movies")
	adminRouter.Use(mw.VerifyToken, mw.AccessGateAdmin)
	adminRouter.POST("", movieHandler.AddMovie)
	adminRouter.PATCH("/:id", movieHandler.EditMovie)
	adminRouter.DELETE("/:id", movieHandler.DeleteMovie)

}
