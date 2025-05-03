package main

import (
	"log"
	"main/internal/routes"
	"main/pkg"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := pkg.Connect()
	if err != nil {
		log.Println(err.Error())
		return
	}

	rdb := pkg.RedisConnect()

	router := routes.InitRoutes(db, rdb)

	router.Static("/profile", "./public/img")

	router.Run("127.0.0.1:8080")
}
