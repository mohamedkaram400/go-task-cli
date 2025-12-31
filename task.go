package main

import (
	"fmt"
	"strconv"
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

func updateTask(taskId string, newDescription string) {
    // Get all tasks
    tasks, err := loadTasks()
    if err != nil {
        fmt.Println("Error loading tasks:", err)
		return
    }

    // change the taskId id from string to int type
    taskIdInt, err := strconv.Atoi(taskId)
    if err != nil {
        fmt.Println("Invalid task ID")
        return
    }
    
    // Loop through tasks and pick the that has the same id
    for i := range tasks {
        if tasks[i].ID == taskIdInt {
            // Update the task description
            tasks[i].Description = newDescription
            break
        }
    }

    if err := saveTasks(tasks); err != nil {
        fmt.Println("Error saving tasks:", err)
    }
}

func markTaskInProgress(taskId string) {
    // Get all tasks
    tasks, err := loadTasks()
    if err != nil {
        fmt.Println("Error loading tasks: ", err)
    }

    // change the taskId id from string to int type
    taskIdInt, err := strconv.Atoi(taskId)

    // Loop through tasks and pick the that has the same id
    for i := range tasks {
        if tasks[i].ID == taskIdInt {
            tasks[i].Status = "in-progress"
            break
        }
    }

    if err := saveTasks(tasks); err != nil {
        fmt.Println("Error saving tasks:", err)
    }
}

func markTaskDone(taskId string) {
    // Get all tasks
    tasks, err := loadTasks()
    if err != nil {
        fmt.Println("Error loading tasks: ", err)
    }

    // change the taskId id from string to int type
    taskIdInt, err := strconv.Atoi(taskId)

    // Loop through tasks and pick the that has the same id
    for i := range tasks {
        if tasks[i].ID == taskIdInt {
            tasks[i].Status = "done"
            break
        }
    }

    if err := saveTasks(tasks); err != nil {
        fmt.Println("Error saving tasks:", err)
    }
}

func markAllTaskDone() {
    // Get all tasks
    tasks, err := loadTasks()
    if err != nil {
        fmt.Println("Error loading tasks: ", err)
    }

    // Loop through tasks 
    for i := range tasks {
        tasks[i].Status = "done"
    }

    if err := saveTasks(tasks); err != nil {
        fmt.Println("Error saving tasks:", err)
    } else {
        fmt.Println("All tasks marked as done")
    }
}

func deleteTask(taskId string) {
    // Get all tasks
    tasks, err := loadTasks()
    if err != nil {
        fmt.Println("Error loading tasks: ", err)
    }

    // change the taskId id from string to int type
    taskIdInt, err := strconv.Atoi(taskId)

    found := false  

    // Loop through tasks and pick the that has the same id
    for i := range tasks {
        if tasks[i].ID == taskIdInt {
            tasks = append(tasks[:i], tasks[i+1:]...)
            found = true
            break
        }
    }

    if !found {
        fmt.Println("Task not found")
        return 
    }

    if err := saveTasks(tasks); err != nil {
        fmt.Println("Error saving tasks:", err)
    } else {
        fmt.Println("Task deleted successfully")
    }
}