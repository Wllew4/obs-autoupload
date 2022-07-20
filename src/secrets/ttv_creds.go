package secrets

import (
	"os"

	"github.com/BurntSushi/toml"
)

var ttv_creds ttv_creds_t
var ttv_cached bool

func TTVCreds() ttv_creds_t {
	if ttv_cached {
		return ttv_creds
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if _, err := toml.DecodeFile(wd+"/.credentials/twitch_secrets.toml", &ttv_creds); err != nil {
		panic(err)
	}

	ttv_cached = true
	return ttv_creds
}

type ttv_creds_t struct {
	Client_id     string
	Client_secret string
}
