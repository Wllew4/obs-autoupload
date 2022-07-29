package secrets

import (
	"auto_upload/src/util"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

var google_creds google_creds_t
var google_cached bool

func GoogleCreds() google_creds_t {
	if google_cached {
		return google_creds
	}

	wd, err := os.Getwd()
	util.CheckErr(err)

	if _, err := toml.DecodeFile(wd+"/.credentials/google_secrets.toml", &google_creds); err != nil {
		panic(err)
	}

	managed := map[string]interface{}{
		"web": map[string]interface{}{
			"client_id":     google_creds.Client_id,
			"client_secret": google_creds.Client_secret,
			"redirect_uris": []string{"http://localhost:8080/"},
			"auth_uri":      "https://accounts.google.com/o/oauth2/auth",
			"token_uri":     "https://accounts.google.com/o/oauth2/token",
		},
	}

	as_json, err := json.Marshal(managed)
	util.CheckErr(err)
	ioutil.WriteFile(wd+"/.credentials/.managed/google_secrets.json", as_json, 0644)

	google_cached = true
	return google_creds
}

type google_creds_t struct {
	Client_id     string
	Client_secret string
}
