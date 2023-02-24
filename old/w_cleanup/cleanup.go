package w_cleanup

import (
	"auto_upload/src/secrets"
	"auto_upload/src/util"
	"auto_upload/src/w_vod"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2/widget"
	"google.golang.org/api/youtube/v3"
)

func add_to_playlist(label *widget.Label, service *youtube.Service, yt_id string) {

	part := []string{
		"snippet",
	}
	playlistItem := youtube.PlaylistItem{
		Snippet: &youtube.PlaylistItemSnippet{
			PlaylistId: secrets.Config().Upload.PLAYLIST_ID,
			ResourceId: &youtube.ResourceId{
				Kind:    "youtube#video",
				VideoId: yt_id,
			},
			Position: 0,
		},
	}
	call := service.PlaylistItems.Insert(part, &playlistItem)
	_, err := call.Do()
	util.CheckErr(err)
	label.SetText("Adding to Playlist: Complete!")
}

func cleanup_step(mode string, label *widget.Label, vod w_vod.VOD, yt_id string) {

	if mode == "move" {
		inputFile, err := os.Open(vod.Path)
		util.CheckErr(err)
		outputFile, err := os.Create(secrets.Config().Files.VOD_ARCHIVE_DIR + filepath.Base(vod.Path))
		util.CheckErr(err)
		defer outputFile.Close()
		_, err = io.Copy(outputFile, inputFile)
		inputFile.Close()
		util.CheckErr(err)

		meta := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", vod.Title, vod.Date, yt_id, vod.Ttv_url, vod.Stream_id)
		ioutil.WriteFile(secrets.Config().Files.VOD_ARCHIVE_DIR+filepath.Base(vod.Path)+".meta", []byte(meta), 0644)
	}
	if mode == "move" || mode == "delete" {
		err := os.Remove(vod.Path)
		util.CheckErr(err)
	}

	label.SetText("Cleanup Step: " + mode + ": Complete!")
}
