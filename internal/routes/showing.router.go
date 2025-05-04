package routes

import (
	"main/internal/handlers"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func initShowingRouter(router *gin.Engine, pg *pgxpool.Pool) {
	showingRepo := repositories.NewShowingRepository(pg)
	showingHandler := handlers.NewShowingHandler(showingRepo)

	showingRouter := router.Group("/showing")
	{
		showingRouter.GET("/schedule", showingHandler.GetMovieSchedules)
		showingRouter.GET("/seat", showingHandler.GetSeatAvailability)
		showingRouter.POST("/seat", showingHandler.ChooseSeat)
	}
}
