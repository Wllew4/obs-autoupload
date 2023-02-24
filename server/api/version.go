package api

import "net/http"

func Version(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("v2.0.0"))
}
