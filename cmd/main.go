package main

import (
	"log"
	"main/internal/routes"
	"main/pkg"

	_ "github.com/joho/godotenv/autoload"
)

// @title 		Tickitz Project
// @version 	1.0
// @description A movie ticket booking application designed for better user experience
// @host 		localhost:8080
// @BasePath 	/
func main() {
	db, err := pkg.Connect()
	if err != nil {
		log.Println(err.Error())
		return
	}

	rdb := pkg.RedisConnect()

	router := routes.InitRoutes(db, rdb)

	router.Static("/public", "./public/img")

	router.Run(":8080")
}
