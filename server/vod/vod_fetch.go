package vod

import (
	"auto_upload/server/secrets"
	"auto_upload/server/ttv"
	"auto_upload/server/util"
	"io/ioutil"
	"path/filepath"
	"time"
)

func FetchVodInfo() VOD {
	last_vod := ttv.TTV_Video()

	return VOD{
		Stream_id: last_vod.Stream_id,
		Title:     last_vod.Title,
		Ttv_url:   last_vod.Url,
		Date:      last_vod.Created_at[:10],
		Path:      findNewestVODFile(),
	}
}

func findNewestVODFile() string {
	cfg, err := secrets.Config()
	util.CheckErr(err)
	dir := cfg.Files.VOD_DIR
	files, err := ioutil.ReadDir(dir)
	util.CheckErr(err)

	var lastTime time.Time
	var lastFile string
	for _, file := range files {
		if file.Mode().IsRegular() {
			if file.ModTime().After(lastTime) {
				lastTime = file.ModTime()
				lastFile = file.Name()
			}
		}
	}
	return filepath.Join(dir, lastFile)
}
