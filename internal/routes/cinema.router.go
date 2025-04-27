package routes

import (
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func initCinemaRouter(router *gin.Engine) {
	cinemaHandler := handlers.NewCinemaHandler()

	cinemaRouter := router.Group("/cinema")
	{
		cinemaRouter.GET("/:schedule_id", cinemaHandler.GetCinema)
	}
}
