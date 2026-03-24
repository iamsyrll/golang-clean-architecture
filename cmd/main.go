package main

import (
	"log"
	"os"

	"golang-clean-arch/config"
	"golang-clean-arch/internal/infrastructure/pgsql"
	"golang-clean-arch/internal/usecase"
	"golang-clean-arch/routes"

	_ "golang-clean-arch/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample API docs auto generate using swagger and gin framework.

//	@contact.name	Clean Architecture API
//	@contact.email	syahrullah.mail@gmail.com

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Description for what is this security definition being used

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

	r.GET("/docs", func(c *gin.Context) {
		c.Redirect(302, "/docs/index.html")
	})

	authorized := r.Group("/docs", gin.BasicAuth(gin.Accounts{
		"admin": "123456",
	}))
	{
		authorized.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	log.Println("Server berjalan pada port :8080")
	r.Run(":8080")
}
