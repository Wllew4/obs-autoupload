package workflow

import (
	"auto_upload/src/secrets"
	"auto_upload/src/ui"
	"auto_upload/src/util"
	"auto_upload/src/yt"
	"fmt"
)

func Start() {
	ui_context := ui.New()
	go step_VODInfo(ui_context)
	ui_context.Window.ShowAndRun()
}

func step_VODInfo(ui_context ui.UIContext) {
	vod_info := fetchVodInfo()
	ui.ShowVOD(ui_context, vod_info, step_Upload)
}

func step_Upload(ui_context ui.UIContext, vod_info util.VOD) {
	fmt.Println("pressed!")
	secrets.GoogleCreds()
	fmt.Println(yt.Upload(
		ui_context,
		vod_info.Path,
		vod_info.Title,
		secrets.Config().Upload.DESCRIPTION,
		secrets.Config().Upload.CATEGORY_ID,
		secrets.Config().Upload.TAGS,
		secrets.Config().Upload.VISIBILITY,
	))
}
