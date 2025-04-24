package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

type Response struct {
	Msg  string `json:"message"`
	Data any    `json:"data"`
}

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
		log.Println("Closing DB...")
		dbClient.Close()
	}()

	// register
	type Users struct {
		Email    string `db:"email" json:"email"`
		Password string `db:"password" json:"-"`
	}

	router.POST("/users", func(ctx *gin.Context) {
		var user Users
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, Response{Msg: "Invalid input"})
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
	type Movies struct {
		Id int `db:"id" json:"id"`
		// Image
		Title        string   `db:"title" json:"movie_title"`
		Genre        string   `db:"movies_genre_id" json:"genre"`
		Duration     int      `db:"duration" json:"duration"`
		Release_date string   `db:"release_date" json:"release"`
		Director     string   `db:"director" json:"director"`
		Casts        []string `db:"casts" json:"casts"`
		Synopsis     string   `db:"synopsis" json:"synopsis"`
	}

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
		query := "SELECT id, title, movies_genre.genre_name, duration, director, casts, synopsis FROM movies JOIN movies_genre ON movies_genre_id = movies_genre.id  WHERE id=$1"
		values := []any{idInt}
		var result Movies
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
		var movie Movies
		if err := ctx.BindJSON(&movie); err != nil {
			ctx.JSON(http.StatusBadRequest, Response{Msg: "Invalid movie"})
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
