package server

import (
	"auto_upload/server/secrets"
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func openInBrowser(path string) {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open", path}
	case "windows":
		args = []string{"cmd", "/c", "start", path}
	default:
		args = []string{"xdg-open", path}
	}
	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Run()
	if err != nil {
		log.Printf("openinbrowser: %v\n", err)
	}
}

func serveConfig(w http.ResponseWriter, r *http.Request) {
	config, err := secrets.Config()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	as_json, err := json.Marshal(config)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(as_json))
}

func Start() {
	mux := http.NewServeMux()

	// frontend
	fileServer := http.FileServer(http.Dir("./client/dist"))
	mux.Handle("/", fileServer)

	// api
	mux.HandleFunc("/api/version", version)
	mux.HandleFunc("/api/verify", verify)
	mux.HandleFunc("/api/config", serveConfig)

	// serve and open
	openInBrowser("http://localhost:80")
	http.ListenAndServe("localhost:80", mux)
}
