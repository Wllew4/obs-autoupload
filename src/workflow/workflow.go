package workflow

import (
	"auto_upload/src/secrets"
	"auto_upload/src/util"
	"auto_upload/src/w_cleanup"
	"auto_upload/src/w_oauth2"
	"auto_upload/src/w_upload"
	"auto_upload/src/w_vod"
	"context"
	"net/http"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type Workflow struct {
	ui_context util.UIContext
	vod_info   w_vod.VOD
	service    *youtube.Service
	yt_id      string
}

func Start() {
	w := Workflow{}
	w.ui_context = util.NewUIContext()
	go w.step_VODInfo()
	w.ui_context.Window.ShowAndRun()
}

func (w Workflow) step_VODInfo() {
	w.vod_info = w_vod.FetchVodInfo()
	w_vod.UI_vod(w.ui_context, w.vod_info, w.step_OAuth2)
}

func (w Workflow) step_OAuth2() {
	secrets.GoogleCreds()
	w_oauth2.GetClient(w.ui_context, w.step_Upload, youtube.YoutubeUploadScope)
}

func (w Workflow) step_Upload(client *http.Client) {

	service, err := youtube.NewService(context.Background(), option.WithHTTPClient(client))
	util.CheckErr(err)
	w.service = service

	w.yt_id = w_upload.Upload(
		w.step_Cleanup,
		w.service,
		w.ui_context,
		w.vod_info.Path,
		w.vod_info.Title,
		w.vod_info.Date,
		w.vod_info.Ttv_url,
		secrets.Config().Upload.DESCRIPTION,
		secrets.Config().Upload.CATEGORY_ID,
		secrets.Config().Upload.TAGS,
		secrets.Config().Upload.VISIBILITY,
	)
}

func (w Workflow) step_Cleanup() {
	w_cleanup.UI_cleanup(w.ui_context, w.service, w.yt_id, w.vod_info)
}
