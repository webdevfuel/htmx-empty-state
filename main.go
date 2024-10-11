package main

import (
	"net/http"

	"github.com/webdevfuel/htmx-empty-state/task"
	"github.com/webdevfuel/htmx-empty-state/template"
)

func main() {
	r := http.NewServeMux()
	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.HandleFunc("GET /tasks", func(w http.ResponseWriter, r *http.Request) {
		tasks := task.ListTasks()
		component := template.EmptyState(tasks)
		component.Render(r.Context(), w)
	})
	r.HandleFunc("DELETE /tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		task.DeleteTask(id)
		tasks := task.ListTasks()
		component := template.Rows(tasks)
		component.Render(r.Context(), w)
	})
	r.HandleFunc("POST /tasks", func(w http.ResponseWriter, r *http.Request) {
		task.AddTask()
		tasks := task.ListTasks()
		component := template.Rows(tasks)
		component.Render(r.Context(), w)
	})
	http.ListenAndServe("localhost:3000", r)
}
