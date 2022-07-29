package workflow

import (
	"auto_upload/src/secrets"
	"auto_upload/src/ui"
	"auto_upload/src/util"
	"auto_upload/src/yt"
	"context"
	"fmt"
	"net/http"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type Workflow struct {
	ui_context ui.UIContext
	vod_info   util.VOD
	service    *youtube.Service
}

func Start() {
	w := Workflow{}
	w.ui_context = ui.New()
	go w.step_VODInfo()
	w.ui_context.Window.ShowAndRun()
}

func (w Workflow) step_VODInfo() {
	w.vod_info = fetchVodInfo()
	ui.ShowVOD(w.ui_context, w.vod_info, w.step_OAuth2)
}

func (w Workflow) step_OAuth2() {
	secrets.GoogleCreds()
	yt.GetClient(w.ui_context, w.step_Upload, youtube.YoutubeUploadScope)
}

func (w Workflow) step_Upload(client *http.Client) {

	service, err := youtube.NewService(context.Background(), option.WithHTTPClient(client))
	util.CheckErr(err)
	w.service = service

	fmt.Println(yt.Upload(
		w.service,
		w.ui_context,
		w.vod_info.Path,
		w.vod_info.Title,
		secrets.Config().Upload.DESCRIPTION,
		secrets.Config().Upload.CATEGORY_ID,
		secrets.Config().Upload.TAGS,
		secrets.Config().Upload.VISIBILITY,
	))
}
