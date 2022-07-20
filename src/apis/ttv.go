package apis

import (
	"auto_upload/src/secrets"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type access_token struct {
	Access_token string
	Expires_in   json.Number
	Token_type   string
}

func AccessToken() string {
	bodyBytes, _ := json.Marshal(map[string]string{
		"client_id":     secrets.TTVCreds().Client_id,
		"client_secret": secrets.TTVCreds().Client_secret,
		"grant_type":    "client_credentials",
	})
	bodyBuffer := bytes.NewBuffer(bodyBytes)
	resp, err := http.Post(
		"https://id.twitch.tv/oauth2/token",
		"application/json",
		bodyBuffer)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	resBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var a access_token
	json.Unmarshal(resBytes, &a)
	return a.Access_token
}
