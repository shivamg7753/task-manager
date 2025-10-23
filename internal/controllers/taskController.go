package controllers

import (
	"net/http"
	"task-manager/internal/config"
	"task-manager/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	config.DB.Create(&task)
	c.JSON(http.StatusOK, task)
}

func GetTasks(c *gin.Context) {
	var tasks []models.Task

	config.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func UpdateTask(c *gin.Context) {
	var task models.Task

	id := c.Param("id")

	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "task not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	config.DB.Save(&task)
	c.JSON(http.StatusOK, task)

}

func DeleteTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")
	if err := config.DB.Delete(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "task not found",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "task deleted",
	})
}
