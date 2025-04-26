package routes

import (
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func initUserRouter(router *gin.Engine) {
	userHandler := handlers.NewUserHandler()

	userRouter := router.Group("/users") 
	{
		userRouter.POST("/register", userHandler.RegistUser)
		userRouter.POST("/login", userHandler.Login)
	}
}
