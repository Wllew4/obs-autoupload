package api

import (
	"auto_upload/server/util"
	"auto_upload/server/vod"
	"encoding/json"
	"net/http"
)

func VOD(w http.ResponseWriter, r *http.Request) {
	vod := vod.FetchVodInfo()
	out, err := json.Marshal(vod)
	util.CheckErr(err)
	w.Write([]byte(out))
}
