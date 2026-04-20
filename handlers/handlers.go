package handlers

import (
	"net/http"
	"taskboard/db"
	"taskboard/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, title, description, status, created_at FROM tasks ORDER BY id DECS")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAT)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, t)
	}

	c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {
	id := c.Param("id")
	var t models.Task

	err := db.DB.QueryRow(
		"SELECT id, title, description, status, created_at FROM tasks WHERE id = $1", id,
	).Scan(&t.ID, &t.Title, &t.Status, &t.Description, &t.CreatedAT)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "задача не найдена"})
		return
	}

	c.JSON(http.StatusOK, t)
}

func CreateTask(c *gin.Context) {
	var t models.Task

	if err := c.ShouldBindJSON(&t); err != nil {
		if t.Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.DB.QueryRow(
		"INSERT INFO tasks (title,description,status) VALUES ($1,$2,$3) RETURING id, created_at",
		t.Title, t.Description, t.Status,
	).Scan(&t.ID, &t.CreatedAT)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, t)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var t models.Task

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.DB.Exec(
		"UPDATE tasks SET title = $1, description = $2,status = $3 WHERE id = $4",
		t.Title, t.Description, t.Status, id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "задача обновлена"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	_, err := db.DB.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "задача удалена"})
}
