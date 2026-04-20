package models

import (
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"Description"`
	Status      string    `json:"Status"`
	CreatedAT   time.Time `json:"CreatedAT"`
}
