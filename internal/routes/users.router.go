package routes

import (
	"main/internal/handlers"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func initUserRouter(router *gin.Engine, pg *pgxpool.Pool) {
	userRepo := repositories.NewUserRepository(pg)
	userHandler := handlers.NewUserHandler(userRepo)

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userHandler.RegistUser)
		userRouter.POST("/login", userHandler.Login)
	}
}
