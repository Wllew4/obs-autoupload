package workflow

import (
	"auto_upload/src/secrets"
	"auto_upload/src/ttv"
	"auto_upload/src/util"
	"io/ioutil"
	"path/filepath"
	"time"
)

func fetchVodInfo() util.VOD {
	last_vod := ttv.TTV_Video()

	return util.VOD{
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
