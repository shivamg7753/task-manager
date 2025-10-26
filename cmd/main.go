package main

import (
	"log"
	"os"
	"task-manager/internal/config"
	"task-manager/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	r := gin.Default()
	config.ConnectDB()
	routes.TaskRoutes(r)
	routes.AuthRoutes(r)
	r.Run(":"+port)
}