package enum

type TaskStatus string 
const (
    TODO TaskStatus = "todo"
    IN_PROGRESS TaskStatus = "in_progress"
    DONE TaskStatus = "done"
)
