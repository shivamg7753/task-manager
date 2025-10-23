package config

import (
	"fmt"
	"task-manager/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnectDB(){
	var err error
	DB,err=gorm.Open(sqlite.Open("task.db"),&gorm.Config{})

	if err!=nil{
		fmt.Println("Failed to connect to database")
		panic(err)
	}

	fmt.Println("connected to database")

	DB.AutoMigrate(&models.Task{})
}