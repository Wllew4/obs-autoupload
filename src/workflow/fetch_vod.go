package workflow

import (
	"auto_upload/src/apis"
	"fmt"
)

type VOD struct {
	title     string
	date      string
	path      string
	stream_id string
	ttv_url   string
	yt_url    string
}

func FetchVodInfo() {
	fmt.Println(apis.AccessToken())
}
