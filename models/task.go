package models

// Task represents a simple task model
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// Simulated database
var tasks = []Task{
	{ID: 1, Title: "Learn Go", Done: false},
	{ID: 2, Title: "Build a web app", Done: false},
}

// GetAllTasks returns all tasks
func GetAllTasks() []Task {
	return tasks
}

// AddTask adds a new task
func AddTask(task Task) {
	task.ID = len(tasks) + 1
	tasks = append(tasks, task)
}
