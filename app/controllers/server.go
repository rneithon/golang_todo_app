package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"../../config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filename ...string) {
	var files []string
	for _, file := range filename {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static", files))

	http.HandleFunc("/", top)
	return http.ListenAndServe(":" + config.Config.Port, nil)
}