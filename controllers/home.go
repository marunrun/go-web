package controllers

import (
	"html/template"
	"net/http"
)

func HomeIndex(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/message.html")
	_ = t.Execute(w, nil)
}
