package api

import (
	"auto_upload/server/secrets"
	"auto_upload/server/util"
	"encoding/json"
	"net/http"
)

func Config(w http.ResponseWriter, r *http.Request) {
	config, err := secrets.Config()
	util.CheckErr(err)
	out, err := json.Marshal(config)
	util.CheckErr(err)
	w.Write([]byte(out))
}
