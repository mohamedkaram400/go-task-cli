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
        // task-cli update 1 "New description"
    case "delete":
        // task-cli delete 1
    case "mark-done":
        // task-cli mark-done 1
    case "mark-in-progress":
        // task-cli mark-in-progress 1
    default:
        fmt.Println("Unknown command:", command)
    }
}