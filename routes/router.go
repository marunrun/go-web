package routes

import (
	"fmt"
	"html"
	"net/http"
)

func fooHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"Hello, %q",html.EscapeString(r.URL.Path))
}



