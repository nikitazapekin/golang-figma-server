package models
 
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
 
var tasks = []Task{
	{ID: 1, Title: "Learn Go", Done: false},
	{ID: 2, Title: "Build a web app", Done: false},
}
 
func GetAllTasks() []Task {
	return tasks
}
 
func AddTask(task Task) {
	task.ID = len(tasks) + 1
	tasks = append(tasks, task)
}
