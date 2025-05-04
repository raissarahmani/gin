package handlers

import (
	"main/internal/models"
	"main/internal/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ShowingHandler struct {
	ShowingRepo *repositories.ShowingRepositories
}

func NewShowingHandler(sr *repositories.ShowingRepositories) *ShowingHandler {
	return &ShowingHandler{
		ShowingRepo: sr,
	}
}

func (s *ShowingHandler) GetMovieSchedules(ctx *gin.Context) {
	movieID, err := strconv.Atoi(ctx.Query("movie_id"))
	if err != nil || movieID <= 0 {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid movie ID"})
		return
	}

	cityID, err := strconv.Atoi(ctx.Query("city_id"))
	if err != nil || cityID <= 0 {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid city ID"})
		return
	}

	cinemaID, err := strconv.Atoi(ctx.Query("cinema_id"))
	if err != nil || cinemaID <= 0 {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid cinema ID"})
		return
	}

	schedules, err := s.ShowingRepo.GetSchedulesByMovieID(ctx.Request.Context(), movieID, cityID, cinemaID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to fetch schedules"})
		return
	}

	if len(schedules) == 0 {
		ctx.JSON(http.StatusNotFound, models.Response{Msg: "No schedule found"})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "Success", Data: schedules})
}

func (s *ShowingHandler) GetSeatAvailability(ctx *gin.Context) {
	movieID, err := strconv.Atoi(ctx.Query("movie_id"))
	if err != nil || movieID <= 0 {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid movie ID"})
		return
	}

	cityID, err := strconv.Atoi(ctx.Query("city_id"))
	if err != nil || cityID <= 0 {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid city ID"})
		return
	}

	cinemaID, err := strconv.Atoi(ctx.Query("cinema_id"))
	if err != nil || cinemaID <= 0 {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid cinema ID"})
		return
	}

	scheduleID, err := strconv.Atoi(ctx.Query("schedule_id"))
	if err != nil || scheduleID <= 0 {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid schedule ID"})
		return
	}

	seats, err := s.ShowingRepo.GetSeatAvailability(ctx.Request.Context(), movieID, cityID, cinemaID, scheduleID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to fetch seat availability"})
		return
	}

	if len(seats) == 0 {
		ctx.JSON(http.StatusNotFound, models.Response{Msg: "No seats found for this schedule"})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "Success", Data: seats})
}

func (s *ShowingHandler) ChooseSeat(ctx *gin.Context) {
	var input models.Schedule
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid input"})
		return
	}

	movieID := input.MovieID
	cityID := input.CityID
	cinemaID := input.CinemaID
	scheduleID := input.ScheduleID
	seatID := input.SeatID

	isAvailable, err := s.ShowingRepo.IsSeatAvailable(ctx.Request.Context(), movieID, cityID, cinemaID, scheduleID, seatID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Error checking seat availability"})
		return
	}
	if !isAvailable {
		ctx.JSON(http.StatusConflict, models.Response{Msg: "Seat already booked"})
		return
	}

	err = s.ShowingRepo.BookSeat(ctx.Request.Context(), seatID, movieID, cityID, cinemaID, scheduleID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to book seat"})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "Seat booked"})
}
