package api

import (
	"auto_upload/server/secrets"
	"net/http"
)

func VerifyConfig(w http.ResponseWriter, r *http.Request) {
	_, err := secrets.Config()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("OK"))
}
