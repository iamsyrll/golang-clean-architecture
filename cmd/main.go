package main

import (
	"log"
	"os"

	"golang-clean-arch/config"
	"golang-clean-arch/internal/infrastructure/pgsql"
	"golang-clean-arch/internal/usecase"
	"golang-clean-arch/routes"

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

	pgdb, err := pgsql.Init()
	if err != nil {
		panic(err)
	}

	pgUserRepo := pgsql.NewUserRepoPG(pgdb)
	authUsecase := usecase.NewAuthUsecase(pgUserRepo, os.Getenv(config.Jwtkey))

	// Initialize Gin router
	// register routes
	r := gin.Default()

	routes.RegisterRoutes(routes.RoutesConfig{
		Router: r,
		AuthUc: authUsecase,
	})

	log.Println("Server berjalan pada port :8080")
	r.Run(":8080")
}
