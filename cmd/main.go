package main

import (
	"log"

	"golang-clean-arch/internal/infrastructure/pgsql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Env Variable failed to load")
	}

	log.Println("ENV loaded sucessfully")

	_, err = pgsql.Init()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	log.Println("Server berjalan pada port :8080")
	r.Run(":8080")
}
