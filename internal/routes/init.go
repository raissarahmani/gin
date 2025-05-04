package routes

import (
	"main/internal/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func InitRoutes(db *pgxpool.Pool, rdb *redis.Client) *gin.Engine {
	router := gin.Default()

	middleware := middlewares.InitMiddleware()

	router.Use(middleware.CORSMiddleware)

	initUserRouter(router, db)
	initMovieRouter(router, db, rdb, middleware)
	initShowingRouter(router, db)
	initProfileRouter(router, middleware)

	return router
}
