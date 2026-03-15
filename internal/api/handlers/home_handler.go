package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func HandleHomeRender(w http.ResponseWriter, r *http.Request) {
	data := new(struct {
		Content string
	})

	data.Content = "Content from server"

	tplPath := filepath.Join("web", "templates", "index.html")
	te, err := template.ParseFiles(tplPath)
	if err != nil {
		http.Error(w, "failed to load template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := te.Execute(w, data); err != nil {
		http.Error(w, "failed to render template", http.StatusInternalServerError)
		return
	}
}
