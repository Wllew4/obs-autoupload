package server

import (
	"auto_upload/server/api"
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

func Start() {
	mux := http.NewServeMux()

	// frontend
	fileServer := http.FileServer(http.Dir("./client/dist"))
	mux.Handle("/", fileServer)

	// api
	mux.HandleFunc("/api/version", api.Version)
	mux.HandleFunc("/api/vod", api.VOD)
	mux.HandleFunc("/api/config", api.Config)
	mux.HandleFunc("/api/upload", api.Upload)

	// serve and open
	openInBrowser("http://localhost:80")
	http.ListenAndServe("localhost:80", mux)
}
