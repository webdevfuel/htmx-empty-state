package main

import (
	"net/http"

	"github.com/webdevfuel/htmx-empty-state/template"
)

func main() {
	r := http.NewServeMux()
	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := template.HelloWorld()
		component.Render(r.Context(), w)
	})
	http.ListenAndServe("localhost:3000", r)
}
