package main

import (
	"github.com/aditya22e/forum-api/api/routes"
	"github.com/aditya22e/forum-api/config"
	"github.com/aditya22e/forum-api/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// Initialize logger
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file: ", err)
	}

	// Initialize database
	if err := config.InitDB(); err != nil {
		logrus.Fatal("Failed to initialize database: ", err)
	}

	// Start scheduler
	schedulerService := services.NewSchedulerService(services.NewTopicService())
	go schedulerService.StartScheduler()
	logrus.Info("Scheduler started")

	// Initialize Gin router
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Setup routes
	routes.SetupRoutes(r)

	// Start server
	port := ":8080"
	logrus.Info("Starting server on ", port)
	if err := r.Run(port); err != nil {
		logrus.Fatal("Failed to start server: ", err)
	}
}
