package routes

import (
	"main/internal/handlers"
	"main/internal/middlewares"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func initOrderRouter(router *gin.Engine, pg *pgxpool.Pool, mdw *middlewares.Middleware) {
	orderRepo := repositories.NewOrderRepository(pg)
	orderHandler := handlers.NewOrderHandler(orderRepo)

	orderRouter := router.Group("/order")
	orderRouter.GET("", mdw.VerifyToken, orderHandler.OrderHistory)
	orderRouter.POST("", mdw.VerifyToken, orderHandler.OrderMovie)
}
