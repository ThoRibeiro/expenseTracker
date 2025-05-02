package main

import (
	"expensetracker/config"
	"expensetracker/database"
	_ "expensetracker/docs"
	"expensetracker/middleware"
	"expensetracker/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title ExpenseTracker API
// @version 1.0
// @description Une API pour gérer vos dépenses quotidiennes.
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Load env & connect DB
	config.Init()
	database.ConnectDatabase()

	r := gin.New()
	r.Use(middleware.Logger()) // logrus hook middleware
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	routes.SetupRoutes(r)

	// Swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := config.GetEnv("PORT", "8080")
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
