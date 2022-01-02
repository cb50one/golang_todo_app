package controllers

import (
	"fmt"
	"golang_todo_app/config"
	"html/template"
	"net/http"
)

func generateHTML(w http.ResponseWriter, data interface{}, filename ...string) {
	var files []string
	for _, file := range filename {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	template := template.Must(template.ParseFiles(files...))
	template.ExecuteTemplate(w, "layout", data)
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
