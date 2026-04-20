package main

import (
	"taskboard/db"
	"taskboard/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	r := gin.Default()

	r.GET("/tasks", handlers.GetTasks)
	r.GET("/tasks/:id", handlers.GetTask)
	r.POST("/tasks", handlers.CreateTask)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)

	r.Run(":8080")
}
