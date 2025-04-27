package main

import (
	"main/internal/repositories"
	"main/internal/routes"
	"main/pkg"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	pkg.Connect()

	repositories.NewUserRepository()
	repositories.NewMovieRepository()
	repositories.NewScheduleRepository()

	router := routes.InitRoutes()

	router.Run("127.0.0.1:8080")
}
