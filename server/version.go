package server

import "net/http"

func version(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("v2.0.0"))
}
