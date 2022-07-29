// ADAPTED FROM: https://github.com/youtube/api-samples/blob/master/go/upload_video.go
package yt

import (
	"auto_upload/src/ui"
	"auto_upload/src/util"
	"fmt"
	"os"
	"strings"

	"fyne.io/fyne/v2/widget"
	"google.golang.org/api/youtube/v3"
)

func Upload(
	service *youtube.Service,
	ui_context ui.UIContext,
	filename string,
	title string,
	description string,
	category string,
	tags []string,
	privacy string,
) string {
	DEBUGNOUPLOAD := false
	keywords := strings.Join(tags, ",")

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

	if !DEBUGNOUPLOAD {
		call := service.Videos.Insert([]string{"snippet", "status"}, upload)

		file, err := os.Open(filename)
		util.CheckErr(err)
		defer file.Close()

		ui_context.SetContent(func() {}, widget.NewLabel("Uploading...\nThis may take awhile..."))

		response, err := call.Media(file).Do()
		util.CheckErr(err)
		fmt.Printf("Upload successful! Video ID: %v\n", response.Id)
		ui_context.SetContent(func() {}, widget.NewLabel("Success!\nhttps://youtube.com/watch?v="+response.Id))
		return response.Id
	}
	return "ID HERE"
}
