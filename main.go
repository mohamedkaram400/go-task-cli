package main

import (
	"fmt"
	"os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide a command")
        return
    }

    command := os.Args[1]

    switch command {

        case "add":
            if len(os.Args) < 3 {
                fmt.Println("Usage: task-cli add \"task description\"")
                return
            }
            addTask(os.Args[2])

        case "list":
            filter := ""
            if len(os.Args) == 3 {
                filter = os.Args[2] 
            }
            listTasks(filter)

        case "update":
            if len(os.Args) != 4 {
                fmt.Println("Usage: task-cli update <id> \"new description\"")
            } 

            updateTask(os.Args[2], os.Args[3])

        case "mark-in-progress":
            if len(os.Args) != 3 {
                fmt.Println("Usage: task-cli mark-in-progress <id>")
            }

            markTaskInProgress(os.Args[2])
        case "mark-done":
            if len(os.Args) != 3 {
                fmt.Println("Usage: task-cli mark-done <id>")
            }

            markTaskDone(os.Args[2])
        case "done":
            if len(os.Args) != 2 {
                fmt.Println("Usage: task-cli done ")
            }

            markAllTaskDone()
        case "delete":
            if len(os.Args) != 3 {
                fmt.Println("Usage: task-cli delete <id>")
            }

            deleteTask(os.Args[2])
        default:
            fmt.Println("Unknown command:", command)
    }
}