package Libs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ReturnJson(code int, msg string, w http.ResponseWriter) {
	Errors := make(map[string]interface{})

	Errors["code"] = code
	Errors["msg"] = msg
	res, _ := json.Marshal(Errors)
	_, _ = fmt.Fprint(w, string(res))

}