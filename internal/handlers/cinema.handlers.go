package handlers

import (
	"main/internal/models"
	"main/internal/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CinemaHandler struct{}

func NewCinemaHandler() *CinemaHandler {
	return &CinemaHandler{}
}

func (c *CinemaHandler) GetCinema(ctx *gin.Context) {
	idParam := ctx.Param("schedule_id")
	schedule_id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "No cinema scheduled"})
		return
	}

	cinemas, err := repositories.CinemaRepo.GetCinemaBySchedule(ctx.Request.Context(), schedule_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Msg:  "Success",
		Data: cinemas,
	})
}
