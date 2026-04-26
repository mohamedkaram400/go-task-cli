package storage

import (
	"encoding/json"
	"os"

	"github.com/mohamedkaram400/go-task-cli/model"
)

const fileName = "storage/tasks.json"

func LoadTasks() ([]model.Task, error) {
    if _, err := os.Stat(fileName); os.IsNotExist(err) {
        return []model.Task{}, nil
    }

    data, err := os.ReadFile(fileName)
    if err != nil {
        return nil, err
    }

    var tasks []model.Task
    err = json.Unmarshal(data, &tasks)
    return tasks, err
}

func SaveTasks(tasks []model.Task) error {
    data, err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(fileName, data, 0644) 
}
