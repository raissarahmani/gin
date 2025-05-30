package handlers

import (
	"log"
	"main/internal/models"
	"main/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderRepo *repositories.OrderRepositories
}

func NewOrderHandler(or *repositories.OrderRepositories) *OrderHandler {
	return &OrderHandler{
		OrderRepo: or,
	}
}

func (o *OrderHandler) OrderMovie(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, models.Response{Msg: "Please login first"})
		return
	}

	var input models.OrderRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid input"})
		return
	}

	// Check seat availability
	seatsAvailable, err := o.OrderRepo.IsSeatAvailable(ctx.Request.Context(), input.MovieID, input.CityID, input.CinemaID, input.ScheduleID, input.SeatIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to check seat availability"})
		return
	}
	if !seatsAvailable {
		ctx.JSON(http.StatusConflict, models.Response{Msg: "Some seats are not available"})
		return
	}

	// Get ticket price
	price, err := o.OrderRepo.GetTicketPrice(ctx.Request.Context(), input.CinemaID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to get price"})
		return
	}
	totalPrice := price * len(input.SeatIDs)

	// Order ticket
	err = o.OrderRepo.BookOrder(ctx.Request.Context(), userID, input, totalPrice)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to order ticket"})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "Order success"})
}

func (o *OrderHandler) OrderHistory(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, models.Response{Msg: "Please login first"})
		return
	}

	orders, err := o.OrderRepo.GetHistory(ctx.Request.Context(), userID)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to get order history"})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Msg:  "Order history retrieved",
		Data: orders,
	})
}
