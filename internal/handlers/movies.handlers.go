package handlers

import (
	"fmt"
	"main/internal/models"
	"main/internal/repositories"
	"net/http"
	"path"
	"strconv"
	"time"

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
	title := ctx.PostForm("title")
	durationStr := ctx.PostForm("duration")
	release := ctx.PostForm("release_date")
	director := ctx.PostForm("director")
	casts := ctx.PostFormArray("casts")
	synopsis := ctx.PostForm("synopsis")

	// input validation
	if title == "" || durationStr == "" || release == "" || director == "" || len(casts) == 0 || synopsis == "" {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "All required fields should be filled"})
		return
	}

	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Duration should be number"})
		return
	}

	releaseDate, err := time.Parse("2006-01-02", release)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid release date format. Use YYYY-MM-DD"})
		return
	}

	// Handle image upload
	file, err := ctx.FormFile("poster")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Image file is required"})
		return
	}
	filename := fmt.Sprintf("%d_poster_%s", time.Now().Unix(), file.Filename)
	filepath := path.Join("public", "img", filename)
	if err := ctx.SaveUploadedFile(file, filepath); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to save poster"})
		return
	}

	// Save to DB
	movie := models.Movies{
		Title:        title,
		Duration:     duration,
		Release_date: releaseDate,
		Director:     director,
		Casts:        casts,
		Synopsis:     synopsis,
		Image:        filename,
	}

	if err := m.MovieRepo.AddNewMovie(ctx.Request.Context(), movie); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to add movie"})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "Movie added successfully"})
}

func (m *MovieHandler) EditMovie(ctx *gin.Context) {
	idStr := ctx.Param("id")
	movieID, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid movie ID"})
		return
	}

	title := ctx.PostForm("title")
	durationStr := ctx.PostForm("duration")
	release := ctx.PostForm("release_date")
	director := ctx.PostForm("director")
	synopsis := ctx.PostForm("synopsis")

	// input validation
	if title == "" || durationStr == "" || release == "" || director == "" || synopsis == "" {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "All required file should be filled"})
		return
	}

	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Duration should be a number"})
		return
	}

	releaseDate, err := time.Parse("2006-01-02", release)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid release date format. Use YYYY-MM-DD"})
		return
	}

	// file upload
	var filename string
	file, err := ctx.FormFile("poster")
	if err == nil {
		filename = fmt.Sprintf("%d_%d_poster_%s", time.Now().Unix(), movieID, file.Filename)
		filepath := path.Join("public", "img", filename)
		if err := ctx.SaveUploadedFile(file, filepath); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to save poster"})
			return
		}
	}

	movie := models.Movies{
		Id:           movieID,
		Title:        title,
		Duration:     duration,
		Release_date: releaseDate,
		Director:     director,
		Synopsis:     synopsis,
	}

	if filename != "" {
		movie.Image = filename
	}

	if err := m.MovieRepo.EditMovie(ctx.Request.Context(), movieID, movie); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Failed to edit movie"})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "Movie updated successfully"})
}

func (m *MovieHandler) DeleteMovie(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
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
