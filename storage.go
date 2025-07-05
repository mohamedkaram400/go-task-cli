package main

import (
    "encoding/json"
    "os"
)

const fileName = "tasks.json"

func loadTasks() ([]Task, error) {
    if _, err := os.Stat(fileName); os.IsNotExist(err) {
        return []Task{}, nil
    }

    data, err := os.ReadFile(fileName)
    if err != nil {
        return nil, err
    }

    var tasks []Task
    err = json.Unmarshal(data, &tasks)
    return tasks, err
}

func saveTasks(tasks []Task) error {
    data, err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(fileName, data, 0644) 
}
