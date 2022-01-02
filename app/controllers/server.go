package controllers

import (
	"fmt"
	"golang_todo_app/app/models"
	"golang_todo_app/config"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func generateHTML(w http.ResponseWriter, data interface{}, filename ...string) {
	var files []string
	for _, file := range filename {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	template := template.Must(template.ParseFiles(files...))
	template.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+$)")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// /todos/edit/1
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, qi)
	}
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))
	http.HandleFunc("/", top)

	port := os.Getenv("PORT")
	return http.ListenAndServe(":"+port, nil)

	// For sqlite
	// return http.ListenAndServe(":"+config.Config.Port, nil)
}
