package main

import (
	"fmt"
	"time"
)

type Task struct {
    ID          int       `json:"id"`
    Description string    `json:"description"`
    Status      string    `json:"status"` 
    CreatedAt   time.Time `json:"createdAt"`
    UpdatedAt   time.Time `json:"updatedAt"`
}

func addTask(description string) {
    tasks, _ := loadTasks()
    newID := 1
    if len(tasks) > 0 {
        newID = tasks[len(tasks)-1].ID + 1
    }

    newTask := Task{
        ID:          newID,
        Description: description,
        Status:      "todo",
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }

    tasks = append(tasks, newTask)

    err := saveTasks(tasks)
    if err != nil {
        fmt.Println("Error saving task:", err)
    } else {
        fmt.Printf("Task added successfully (ID: %d)\n", newID)
    }
}

func listTasks(filter string) {
    tasks, err := loadTasks()
    if err != nil {
        fmt.Println("Error loading tasks:", err)
        return
    }

    for _, task := range tasks {
        if filter == "" || task.Status == filter {
            fmt.Printf("ID: %d | %s | Status: %s | Updated: %s\n",
                task.ID, task.Description, task.Status, task.UpdatedAt.Format(time.RFC1123))
        }
    }
}