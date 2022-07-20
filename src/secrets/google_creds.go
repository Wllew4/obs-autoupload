package secrets

import (
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
	if err != nil {
		panic(err)
	}

	if _, err := toml.DecodeFile(wd+"/.credentials/twitch_secrets.toml", &google_creds); err != nil {
		panic(err)
	}

	type managed_creds_inner struct {
		client_id       string
		client_secret   string
		redirect_uris   string
		auth_uri        string
		token_uristring string
	}

	type managed_creds struct {
		web managed_creds_inner
	}

	google_cached = true
	return google_creds
}

type google_creds_t struct {
	Client_id     string
	Client_secret string
}
