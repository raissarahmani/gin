package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
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
		log.Printf("Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		log.Println("Closing DB...")
		dbClient.Close()
	}()
}