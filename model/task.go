package model

import (
	"time"

	enum "github.com/mohamedkaram400/go-task-cli/Enum"
)


type Task struct {
	ID          int       `json:"id"`
	Title       string       `json:"title"`
	Description string    `json:"description"`
	Status      enum.TaskStatus    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}