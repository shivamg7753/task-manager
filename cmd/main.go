package main

import (
	"task-manager/internal/config"
	"task-manager/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r:= gin.Default()
	config.ConnectDB()
	routes.TaskRoutes(r)
  r.Run(":8080")
}
