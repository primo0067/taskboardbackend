package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"taskboard/models"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения: ", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("БД не доступна:", err)
	}

	createTable()
	log.Println("БД подключена")
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS tasks(
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		status TEXT DEFAULT 'todo',
		createdAT TIMESTAMP DEFAULT NOW()
	);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Ошибка создания таблицы: ", err)
	}
}

func GetsDB() ([]models.Task, error) {
	rows, err := DB.Query("SELECT id, title, description, status, createdAT FROM tasks ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAT)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func GetDB(id string) (models.Task, error) {
	var t models.Task

	err := DB.QueryRow(
		"SELECT id, title, description, status, createdAT FROM tasks WHERE id = $1", id,
	).Scan(&t.ID, &t.Title, &t.Status, &t.Description, &t.CreatedAT)

	if err != nil {
		return t, err
	}
	return t, err
}

func CreateDB(t *models.Task) error {

	if t.Title == "" {
		return errors.New("задача пуста")
	}
	err := DB.QueryRow(
		"INSERT INTO tasks (title,description,status) VALUES ($1,$2,$3) RETURNING id, createdAT",
		t.Title, t.Description, t.Status,
	).Scan(&t.ID, &t.CreatedAT)
	if err != nil {
		return err
	}
	return nil
}

func UpdateDB(id string, t *models.Task) error {
	_, err := DB.Exec(
		"UPDATE tasks SET title = $1, description = $2,status = $3 WHERE id = $4",
		t.Title, t.Description, t.Status, id,
	)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDB(id string) error {
	_, err := DB.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
