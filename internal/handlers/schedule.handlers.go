package handlers

import (
	"main/internal/models"
	"main/internal/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ScheduleHandler struct{}

func NewScheduleHandler() *ScheduleHandler {
	return &ScheduleHandler{}
}

func (s *ScheduleHandler) GetMovieSchedules(ctx *gin.Context) {
	idParam := ctx.Param("movie_id")
	movie_id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "No schedule available for this movie"})
		return
	}

	schedules, err := repositories.ScheduleRepo.GetSchedulesByMovieID(ctx.Request.Context(), movie_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Msg:  "Success",
		Data: schedules,
	})
}
