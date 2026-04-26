package worker

import (
	"fmt"
	"sync"

	"github.com/mohamedkaram400/go-task-cli/model"
	"github.com/mohamedkaram400/go-task-cli/storage"
)


var mu sync.Mutex

func StartTaskManager(id int, cmdChan chan model.Command) {
	for cmd := range cmdChan {
		fmt.Println("Worker", id, "processing", cmd.Action)

		switch cmd.Action {

		case "add":
			mu.Lock()
			
			data := cmd.Payload.(model.Task)

			tasks, _ := storage.LoadTasks()

			newID := 1
			if len(tasks) > 0 {
				newID = tasks[len(tasks)-1].ID + 1
			}

			data.ID = newID
			tasks = append(tasks, data)
			storage.SaveTasks(tasks)
			mu.Unlock()

			cmd.Result <- newID

		case "list":
			tasks, _ := storage.LoadTasks()
			cmd.Result <- tasks

		case "delete":
			id := cmd.Payload.(int)
			tasks, _ := storage.LoadTasks()

			for i := range tasks {
				if tasks[i].ID == id {
					tasks = append(tasks[:i], tasks[i+1:]...)
					break
				}
			}

			storage.SaveTasks(tasks)
			cmd.Result <- "deleted"

		}
	}
}