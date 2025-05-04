package routes

import (
	"main/internal/handlers"
	"main/internal/middlewares"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func initProfileRouter(router *gin.Engine, pg *pgxpool.Pool, mdw *middlewares.Middleware) {
	profileRepo := repositories.NewProfileRepository(pg)
	profileHandler := handlers.NewProfileHandler(profileRepo)

	profileRouter := router.Group("/profile")
	profileRouter.GET("", mdw.VerifyToken, profileHandler.GetProfile)
	profileRouter.PATCH("/edit", mdw.VerifyToken, profileHandler.EditProfile)
	profileRouter.PATCH("/pass", mdw.VerifyToken, profileHandler.ChangePassword)
}
