package routes

import (
	"task-manager/internal/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine){
	auth :=r.Group("/auth")
	{
		auth.POST("/register",controllers.Register)
		auth.POST("/login",controllers.Login)
	}
}