package workflow

import (
	"auto_upload/src/secrets"
	"auto_upload/src/ui"
	"auto_upload/src/util"
	"auto_upload/src/yt"
	"fmt"
)

func Start() {
	context := ui.New()
	go step_VODInfo(context)
	context.Window.ShowAndRun()
}

func step_VODInfo(context ui.UIContext) {
	vod_info := fetchVodInfo()
	ui.ShowVOD(vod_info, step_Upload, context)
}

func step_Upload(context ui.UIContext, vod_info util.VOD) {
	fmt.Println("pressed!")
	secrets.GoogleCreds()
	yt.Upload(
		vod_info.Path,
		vod_info.Title,
		secrets.Config().Upload.DESCRIPTION,
		secrets.Config().Upload.CATEGORY_ID,
		secrets.Config().Upload.TAGS,
		secrets.Config().Upload.VISIBILITY,
	)
}
