package interfaces

import "github.com/mohamedkaram400/go-task-cli/model"

type TaskRepository interface {
	Add(Task model.Task) error
	List() ([]model.Task, error)
	Delete(id int) error
}