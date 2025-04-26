package main

import (
	"context"
	"fmt"
	"log"
	"main/internal/models"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	router := gin.Default()

	dbEnv := []any{}
	dbEnv = append(dbEnv, os.Getenv("DBUSER"))
	dbEnv = append(dbEnv, os.Getenv("DBPASS"))
	dbEnv = append(dbEnv, os.Getenv("DBHOST"))
	dbEnv = append(dbEnv, os.Getenv("DBPORT"))
	dbEnv = append(dbEnv, os.Getenv("DBNAME"))
	dbString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbEnv...)
	dbClient, err := pgxpool.New(context.Background(), dbString)
	if err != nil {
		log.Printf("Failed to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		dbClient.Close()
	}()

	// register
	router.POST("/users", func(ctx *gin.Context) {
		var user models.Users
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid input"})
			return
		}
		query := "INSERT INTO users (email, password) VALUES ($1, $2)"
		values := []any{}
		cmd, err := dbClient.Exec(ctx.Request.Context(), query, values...)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "Internal server error",
			})
			return
		}
		if cmd.RowsAffected() == 0 {
			log.Println("Register failed")
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	})

	// movies
	router.GET("/movies/:id", func(ctx *gin.Context) {
		idStr, ok := ctx.Params.Get("id")
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "Param id is needed",
			})
			return
		}
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "Internal server error",
			})
			return
		}
		query := "SELECT m.id, m.title, m.movies_genre_id, m.duration, m.director, m.casts, m.synopsis FROM movies m  WHERE id=$1"
		values := []any{idInt}
		var result models.Movies
		if err := dbClient.QueryRow(context.Background(), query, values...).Scan(&result.Id, &result.Title, &result.Genre, &result.Duration, &result.Director, &result.Casts, &result.Synopsis); err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "System error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"Msg":    "success",
			"movies": result,
		})
	})

	router.POST("/movies", func(ctx *gin.Context) {
		var movie models.Movies
		if err := ctx.BindJSON(&movie); err != nil {
			ctx.JSON(http.StatusBadRequest, models.Response{Msg: "Invalid movie"})
			return
		}
		query := "INSERT INTO movies (title, director, casts) VALUES ($1, $2, $3)"
		values := []any{}
		cmd, err := dbClient.Exec(ctx.Request.Context(), query, values...)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "Internal server error",
			})
			return
		}
		if cmd.RowsAffected() == 0 {
			log.Println("Add movie failed")
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	})

	router.Run("127.0.0.1:8080")
}
