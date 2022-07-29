package w_vod

import (
	"auto_upload/src/api_ttv"
	"auto_upload/src/secrets"
	"auto_upload/src/util"
	"io/ioutil"
	"path/filepath"
	"time"
)

func FetchVodInfo() VOD {
	last_vod := api_ttv.TTV_Video()

	return VOD{
		Stream_id: last_vod.Stream_id,
		Title:     last_vod.Title,
		Ttv_url:   last_vod.Url,
		Date:      last_vod.Created_at[:10],
		Path:      findNewestVODFile(),
	}
}

func findNewestVODFile() string {
	dir := secrets.Config().Files.VOD_DIR
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
