package routes

import (
	"task-manager/internal/controllers"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine){
	task:=r.Group("/tasks")

	{
		task.POST("/",controllers.CreateTask)
		task.GET("/",controllers.GetTasks)
		task.PUT("/:id",controllers.UpdateTask)
		task.DELETE("/:id",controllers.DeleteTask)
	}
}