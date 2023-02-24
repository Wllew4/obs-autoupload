package api

import (
	"fmt"
	"net/http"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
}
