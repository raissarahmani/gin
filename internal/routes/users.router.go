package routes

import (
	"main/internal/handlers"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func initUserRouter(router *gin.Engine, pg *pgxpool.Pool, rdc *redis.Client) {
	userRepo := repositories.NewUserRepository(pg, rdc)
	userHandler := handlers.NewUserHandler(userRepo)

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userHandler.RegistUser)
		userRouter.POST("/login", userHandler.Login)
	}
}
