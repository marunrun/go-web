package controllers

import (
	"fmt"
	"net/http"
)

func HomeIndex(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	_, _ = fmt.Fprintln(w, "Welcome to my Home !")
}
