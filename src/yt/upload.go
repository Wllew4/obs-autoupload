// ADAPTED FROM: https://github.com/youtube/api-samples/blob/master/go/upload_video.go
package yt

import (
	"auto_upload/src/ui"
	"auto_upload/src/util"
	"os"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func Upload(
	ui_context ui.UIContext,
	filename string,
	title string,
	description string,
	category string,
	tags []string,
	privacy string,
) string {
	keywords := strings.Join(tags, ",")

	client := getClient(ui_context, youtube.YoutubeUploadScope)

	service, err := youtube.NewService(context.Background(), option.WithHTTPClient(client))
	util.CheckErr(err)

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

	// call := service.Videos.Insert([]string{"snippet", "status"}, upload)

	file, err := os.Open(filename)
	util.CheckErr(err)
	defer file.Close()

	// response, err := call.Media(file).Do()
	util.CheckErr(err)
	// fmt.Printf("Upload successful! Video ID: %v\n", response.Id)
	// return response.Id
	_ = service
	_ = upload
	return "ID HERE"
}
