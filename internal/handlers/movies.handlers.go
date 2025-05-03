package handlers

import (
	"main/internal/models"
	"main/internal/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	MovieRepo *repositories.MovieRepositories
}

func NewMovieHandler(mr *repositories.MovieRepositories) *MovieHandler {
	return &MovieHandler{
		MovieRepo: mr,
	}
}

func (m *MovieHandler) AllMovies(ctx *gin.Context) {
	pageParam := ctx.DefaultQuery("page", "1") // from ?page=2
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	const pageSize = 5
	offset := (page - 1) * pageSize

	movies, err := m.MovieRepo.ShowAllMovies(ctx.Request.Context(), pageSize, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "Success", Data: movies})
}

func (m *MovieHandler) MovieDetail(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid movie"})
		return
	}

	detail, err := m.MovieRepo.ShowMovieDetail(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "Success", Data: detail})
}

func (m *MovieHandler) FilterMovies(ctx *gin.Context) {
	title := ctx.DefaultQuery("title", "")
	genre := ctx.DefaultQuery("genre", "")

	var movies []models.Movies
	var err error
	var limit = 5
	var page = 1
	var offset = (page - 1) * limit

	if title != "" && genre != "" {
		movies, err = m.MovieRepo.FilterMoviesByTitleAndGenre(ctx.Request.Context(), title, genre)
	} else if title != "" {
		movies, err = m.MovieRepo.FilterMoviesByTitle(ctx.Request.Context(), title)
	} else if genre != "" {
		movies, err = m.MovieRepo.FilterMoviesByGenre(ctx.Request.Context(), genre)
	} else {
		movies, err = m.MovieRepo.ShowAllMovies(ctx.Request.Context(), limit, offset)
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "Success", Data: movies})
}

func (m *MovieHandler) UpcomingMovies(ctx *gin.Context) {
	movies, err := m.MovieRepo.ShowUpcomingMovies(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{Msg: "Success", Data: movies})
}

func (m *MovieHandler) AddMovie(ctx *gin.Context) {
	var newMovie models.Movies
	if err := ctx.ShouldBindJSON(&newMovie); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid request body"})
		return
	}
	err := m.MovieRepo.AddNewMovie(ctx.Request.Context(), newMovie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to add movie"})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{Msg: "Movie added successfully"})
}

func (m *MovieHandler) EditMovie(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid movie ID"})
		return
	}

	var editedMovie models.Movies
	if err := ctx.ShouldBindJSON(&editedMovie); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid request body"})
		return
	}

	err = m.MovieRepo.EditMovie(ctx.Request.Context(), id, editedMovie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Edit movie failed"})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "Movie updated"})
}

func (m *MovieHandler) DeleteMovie(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid movie ID"})
		return
	}

	err = m.MovieRepo.DeleteMovie(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Delete movie failed"})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "Movie deleted"})
}
