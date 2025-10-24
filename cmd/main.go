package main

import (
	"log"
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
	r := gin.Default()
	config.ConnectDB()
	routes.TaskRoutes(r)
	routes.AuthRoutes(r)
	r.Run(":8080")
}
