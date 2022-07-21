package workflow

import (
	"auto_upload/src/apis"
)

type VOD struct {
	Title     string
	Date      string
	Path      string
	Stream_id string
	Ttv_url   string
	Yt_url    string
}

func FetchVodInfo() VOD {
	last_vod := apis.TTV_Video()

	return VOD{
		Stream_id: last_vod.Stream_id,
		Title:     last_vod.Title,
		Ttv_url:   last_vod.Url,
		Date:      last_vod.Created_at[:10],
	}
}
