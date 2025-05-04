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
