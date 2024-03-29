// ADAPTED FROM: https://github.com/youtube/api-samples/blob/master/go/upload_video.go
package w_upload

import (
	"auto_upload/src/util"
	"fmt"
	"os"
	"strings"

	"google.golang.org/api/youtube/v3"
)

func Upload(
	nextStep func(string),
	service *youtube.Service,
	ui_context util.UIContext,
	filename string,
	title string,
	date string,
	ttv_url string,
	description string,
	category string,
	tags []string,
	privacy string,
) {
	DEBUGNOUPLOAD := false
	keywords := strings.Join(tags, ",")

	description = strings.Replace(description, "{TITLE}", title, -1)
	description = strings.Replace(description, "{DATE}", date, -1)
	description = strings.Replace(description, "{URL}", ttv_url, -1)

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       title,
			Description: description,
			CategoryId:  category,
		},
		Status: &youtube.VideoStatus{PrivacyStatus: privacy},
	}

	// The API returns a 400 Bad Request response if tags is an empty string.
	if strings.Trim(keywords, "") != "" {
		upload.Snippet.Tags = strings.Split(keywords, ",")
	}

	ui_uploading(ui_context)
	var id string = "ID-HERE"
	if !DEBUGNOUPLOAD {
		call := service.Videos.Insert([]string{"snippet", "status"}, upload)

		file, err := os.Open(filename)
		util.CheckErr(err)
		defer file.Close()

		response, err := call.Media(file).Do()
		util.CheckErr(err)
		fmt.Printf("Upload successful! Video ID: %v\n", response.Id)
		id = response.Id
	}

	ui_uploaded(ui_context, nextStep, id)
}
