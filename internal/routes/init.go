package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	initUserRouter(router)
	initMovieRouter(router)
	initScheduleRouter(router)

	return router
}
