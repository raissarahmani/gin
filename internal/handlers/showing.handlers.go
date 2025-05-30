package handlers

import (
	"log"
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

func (s *ShowingHandler) GetFilteredShowings(ctx *gin.Context) {
	movieIDStr := ctx.Query("movie_id")
	cityIDStr := ctx.Query("city_id")
	scheduleIDStr := ctx.Query("schedule_id")
	cinemaIDStr := ctx.Query("cinema_id")
	var result = make(map[string]interface{})

	movieID, err := strconv.Atoi(movieIDStr)
	if err != nil || movieID <= 0 {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid or missing movie_id"})
		return
	}

	cityID, err := strconv.Atoi(cityIDStr)
	scheduleID, err := strconv.Atoi(scheduleIDStr)
	cinemaID, err := strconv.Atoi(cinemaIDStr)

	// Get all schedule and all city for movie showings
	if cityIDStr == "" && scheduleIDStr == "" {
		schedules, err := s.ShowingRepo.GetSchedulesByMovie(ctx.Request.Context(), movieID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to fetch schedules"})
			return
		}

		cities, err := s.ShowingRepo.GetCitiesByMovie(ctx.Request.Context(), movieID)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to fetch cities"})
			return
		}

		result["schedules"] = schedules
		result["cities"] = cities
	}

	// Get cinema showing movie in selected schedule and city
	if cityIDStr != "" && scheduleIDStr != "" && cinemaIDStr == "" {
		cinemas, err := s.ShowingRepo.GetCinemasByFilters(ctx.Request.Context(), movieID, cityID, scheduleID)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to fetch cinemas"})
			return
		}
		result["cinemas"] = cinemas
	}

	// Book schedule
	if cityIDStr != "" && scheduleIDStr != "" && cinemaIDStr != "" {
		booking, err := s.ShowingRepo.BookSchedule(ctx.Request.Context(), movieID, cityID, scheduleID, cinemaID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to book schedule"})
			return
		}
		result["book"] = booking
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "Success", Data: result})
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
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to fetch seat availability"})
		return
	}

	if len(seats) == 0 {
		ctx.JSON(http.StatusNotFound, models.Response{Msg: "No seats found for this schedule"})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "Success", Data: seats})
}
