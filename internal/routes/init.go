package routes

import (
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	repositories.NewMovieRepository()
	repositories.NewUserRepository()

	initUserRouter(router)
	initMovieRouter(router)

	return router
}