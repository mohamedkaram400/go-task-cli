package main

import (
	"fmt"
	"os"
	"time"

	enum "github.com/mohamedkaram400/go-task-cli/Enum"
	"github.com/mohamedkaram400/go-task-cli/model"
	"github.com/mohamedkaram400/go-task-cli/worker"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide a command")
        return
    }

    command := os.Args[1]

	cmdChan := make(chan model.Command)
    workerCount := 3

	// start goroutine
    for i := 1; i <= workerCount; i++ {
        go worker.StartTaskManager(i, cmdChan)
    }

	resultChan := make(chan interface{})

	switch command {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: add \"title\" \"description\"")
			return
		}

		cmdChan <- model.Command{
			Action: "add",
			Payload: model.Task{
				Title: os.Args[2],
				Description: os.Args[3],
				Status: enum.TODO,
			},
			Result: resultChan,
		}

		result := <-resultChan
		fmt.Println("Task added with ID:", result)

	case "list":
		cmdChan <- model.Command{
			Action: "list",
			Result: resultChan,
		}

		tasks := (<-resultChan).([]model.Task)

		for _, task := range tasks {
            fmt.Printf("ID: %d | Title: %s |Description: %s | Status: %s | Updated: %s\n",
				task.ID, task.Title, task.Description, task.Status, task.UpdatedAt.Format(time.RFC1123))
		}

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: delete <id>")
			return
		}

		id := 0
		fmt.Sscanf(os.Args[2], "%d", &id)

		cmdChan <- model.Command{
			Action:  "delete",
			Payload: id,
			Result:  resultChan,
		}

		fmt.Println(<-resultChan)
	}

    // command := os.Args[1]

    // switch command {

    //     case "add":
    //         if len(os.Args) < 3 {
    //             fmt.Println("Usage: task-cli add \"task description\"")
    //             return
    //         }
    //         logic.AddTask(os.Args[2], os.Args[3])

    //     case "list":
    //         filter := ""
    //         if len(os.Args) == 3 {
    //             filter = os.Args[2] 
    //         }
    //         logic.ListTasks(filter)

    //     case "update":
    //         if len(os.Args) != 4 {
    //             fmt.Println("Usage: task-cli update <id> \"new description\"")
    //         } 
    //         logic.UpdateTask(os.Args[2], os.Args[3])

    //     case "mark-in-progress":
    //         if len(os.Args) != 3 {
    //             fmt.Println("Usage: task-cli mark-in-progress <id>")
    //         }
    //         logic.MarkTaskInProgress(os.Args[2])

    //     case "mark-done":
    //         if len(os.Args) != 3 {
    //             fmt.Println("Usage: task-cli mark-done <id>")
    //         }
    //         logic.MarkTaskDone(os.Args[2])

    //     case "done":
    //         if len(os.Args) != 2 {
    //             fmt.Println("Usage: task-cli done ")
    //         }
    //         logic.MarkAllTaskDone()

    //     case "delete":
    //         if len(os.Args) != 3 {
    //             fmt.Println("Usage: task-cli delete <id>")
    //         }
    //         logic.DeleteTask(os.Args[2])
            
    //     default:
    //         fmt.Println("Unknown command:", command)
    // }
}