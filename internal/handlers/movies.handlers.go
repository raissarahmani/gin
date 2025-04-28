package handlers

import (
	"main/internal/models"
	"main/internal/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct{}

func NewMovieHandler() *MovieHandler {
	return &MovieHandler{}
}

func (m *MovieHandler) AllMovies(ctx *gin.Context) {
	movies, err := repositories.MovieRepo.ShowAllMovies(ctx.Request.Context())
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

	detail, err := repositories.MovieRepo.ShowMovieDetail(ctx.Request.Context(), id)
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

	if title != "" && genre != "" {
		movies, err = repositories.MovieRepo.FilterMoviesByTitleAndGenre(ctx.Request.Context(), title, genre)
	} else if title != "" {
		movies, err = repositories.MovieRepo.FilterMoviesByTitle(ctx.Request.Context(), title)
	} else if genre != "" {
		movies, err = repositories.MovieRepo.FilterMoviesByGenre(ctx.Request.Context(), genre)
	} else {
		movies, err = repositories.MovieRepo.ShowAllMovies(ctx.Request.Context())
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "Success", Data: movies})
}

func (m *MovieHandler) UpcomingMovies(ctx *gin.Context) {
	movies, err := repositories.MovieRepo.ShowUpcomingMovies(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Msg: "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{Msg: "Success", Data: movies})
}

// func (m *MovieHandler) AddMovie(ctx *gin.Context) {
// 	query := "INSERT INTO movies (title, director, casts) VALUES ($1, $2, $3)"
// 	values := []any{}
// 	cmd, err := pkg.DB.Exec(ctx.Request.Context(), query, values...)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"msg": "Terjadi kesalahan server",
// 		})
// 		return
// 	}
// 	if cmd.RowsAffected() == 0 {
// 		log.Println("Query Gagal, Tidak merubah data di DB")
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg": "success",
// 	})
// }
