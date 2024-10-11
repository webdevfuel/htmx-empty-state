package task

import (
	"fmt"

	"github.com/google/uuid"
)

type Task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var tasks []Task = []Task{
	{
		ID:   "1ca9eab4-c23b-45a1-9abb-6a8ec1e46baa",
		Name: "Task 1",
	},
	{
		ID:   "d03e6daa-3cad-4c8e-97d6-2e2fddfc9f8f",
		Name: "Task 2",
	},
	{
		ID:   "456fc3b7-03a1-4642-9a5d-80d99a93c013",
		Name: "Task 3",
	},
	{
		ID:   "1dfe9b83-0068-4efc-9160-9a632aa96cfa",
		Name: "Task 4",
	},
}

func ListTasks() []Task {
	return tasks
}

func AddTask() Task {
	uuid, _ := uuid.NewRandom()
	task := Task{
		ID:   uuid.String(),
		Name: fmt.Sprintf("Task %d", len(tasks)+1),
	}
	tasks = append(tasks, task)
	return task
}

func DeleteTask(id string) {
	var filteredTasks []Task
	for _, task := range tasks {
		if task.ID != id {
			filteredTasks = append(filteredTasks, task)
		}
	}
	tasks = filteredTasks
}
