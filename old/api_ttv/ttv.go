package api_ttv

import (
	"auto_upload/src/secrets"
	"auto_upload/src/util"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func TTV_AccessToken() ttv_access_token {
	bodyBytes, _ := json.Marshal(map[string]string{
		"client_id":     secrets.TTVCreds().Client_id,
		"client_secret": secrets.TTVCreds().Client_secret,
		"grant_type":    "client_credentials",
	})
	bodyBuffer := bytes.NewBuffer(bodyBytes)
	res, err := http.Post(
		"https://id.twitch.tv/oauth2/token",
		"application/json",
		bodyBuffer,
	)
	util.CheckErr(err)
	defer res.Body.Close()

	resBytes, err := ioutil.ReadAll(res.Body)
	util.CheckErr(err)

	var access_token ttv_access_token
	json.Unmarshal(resBytes, &access_token)
	return access_token
}

func TTV_User() ttv_user {
	client := http.Client{}
	config := secrets.Config()
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"https://api.twitch.tv/helix/users?login=%s",
			config.TTV.LOGIN,
		),
		nil,
	)
	util.CheckErr(err)

	req.Header.Set("Authorization", "Bearer "+TTV_AccessToken().Access_token)
	req.Header.Set("Client-Id", secrets.TTVCreds().Client_id)
	res, err := client.Do(req)
	util.CheckErr(err)
	defer res.Body.Close()

	resBytes, err := ioutil.ReadAll(res.Body)
	util.CheckErr(err)

	var user ttv_res[ttv_user]
	err = json.Unmarshal(resBytes, &user)
	util.CheckErr(err)
	return user.Data[0]
}

func TTV_Video() ttv_video {
	client := http.Client{}
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"https://api.twitch.tv/helix/videos?user_id=%s&first=1&sort=time&type=archive",
			TTV_User().Id,
		),
		nil,
	)
	util.CheckErr(err)

	req.Header.Set("Authorization", "Bearer "+TTV_AccessToken().Access_token)
	req.Header.Set("Client-Id", secrets.TTVCreds().Client_id)
	res, err := client.Do(req)
	util.CheckErr(err)
	defer res.Body.Close()

	resBytes, err := ioutil.ReadAll(res.Body)
	util.CheckErr(err)

	var video ttv_res[ttv_video]
	err = json.Unmarshal(resBytes, &video)
	util.CheckErr(err)
	return video.Data[0]
}
