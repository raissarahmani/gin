package routes

import (
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func initScheduleRouter(router *gin.Engine) {
	scheduleHandler := handlers.NewScheduleHandler()

	scheduleRouter := router.Group("/schedule")
	{
		scheduleRouter.GET("/:movie_id", scheduleHandler.GetMovieSchedules)
	}
}
