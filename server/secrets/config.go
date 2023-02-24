package secrets

import (
	"os"

	"github.com/BurntSushi/toml"
)

var config config_t
var config_cached bool

func Config() (*config_t, error) {
	if config_cached {
		return &config, nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	if _, err := toml.DecodeFile(wd+"/devconfig.toml", &config); err != nil {
		return nil, err
	}

	config_cached = true
	return &config, nil
}

type config_t struct {
	Files  config_files
	TTV    config_ttv
	Upload config_upload
}

type config_files struct {
	VOD_DIR          string
	POST_UPLOAD_STEP string
	VOD_ARCHIVE_DIR  string
}

type config_ttv struct {
	LOGIN string
}

type config_upload struct {
	VISIBILITY  string
	DESCRIPTION string
	CATEGORY_ID string
	TAGS        []string
	PLAYLIST_ID string
}
