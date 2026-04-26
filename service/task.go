package service

import (
	"fmt"
	"strconv"
	"time"

	enum "github.com/mohamedkaram400/go-task-cli/Enum"
	"github.com/mohamedkaram400/go-task-cli/model"
	"github.com/mohamedkaram400/go-task-cli/storage"
)


func AddTask(title string, description string) {

	tasks, _ := storage.LoadTasks()
	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	newTask := model.Task{
		ID:          newID,
		Title:       title,
		Description: description,
		Status:      enum.TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	err := storage.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving task:", err)
	} else {
		fmt.Printf("Task added successfully (ID: %d)\n", newID)
	}
}

func ListTasks(filter string) {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}


	for _, task := range tasks {
		if filter == "" || task.Status == enum.TaskStatus(filter) {
			fmt.Printf("ID: %d | Title: %s |Description: %s | Status: %s | Updated: %s\n",
				task.ID, task.Title, task.Description, task.Status, task.UpdatedAt.Format(time.RFC1123))
		}
	}
}

func UpdateTask(taskId string, newDescription string) {
    // Get all tasks
    tasks, err := storage.LoadTasks()
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

    if err := storage.SaveTasks(tasks); err != nil {
        fmt.Println("Error saving tasks:", err)
    }
}

func MarkTaskInProgress(taskId string) {
    // Get all tasks
    tasks, err := storage.LoadTasks()
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

    if err := storage.SaveTasks(tasks); err != nil {
        fmt.Println("Error saving tasks:", err)
    }
}

func MarkTaskDone(taskId string) {
    // Get all tasks
    tasks, err := storage.LoadTasks()
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

    if err := storage.SaveTasks(tasks); err != nil {
        fmt.Println("Error saving tasks:", err)
    }
}

func MarkAllTaskDone() {
    // Get all tasks
    tasks, err := storage.LoadTasks()
    if err != nil {
        fmt.Println("Error loading tasks: ", err)
    }

    // Loop through tasks 
    for i := range tasks {
        tasks[i].Status = "done"
    }

    if err := storage.SaveTasks(tasks); err != nil {
        fmt.Println("Error saving tasks:", err)
    } else {
        fmt.Println("All tasks marked as done")
    }
}

func DeleteTask(taskId string) {
    // Get all tasks
    tasks, err := storage.LoadTasks()
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

    if err := storage.SaveTasks(tasks); err != nil {
        fmt.Println("Error saving tasks:", err)
    } else {
        fmt.Println("Task deleted successfully")
    }
}