package w_cleanup

import (
	"auto_upload/src/secrets"
	"auto_upload/src/util"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func UI_cleanup(ui_context util.UIContext) {
	content := container.New(layout.NewFormLayout(),
		container.NewVBox(
			widget.NewLabel("Post-Upload Step:     \t"),
			widget.NewLabel("VOD Archive Directory:\t"),
			widget.NewLabel("Playlist ID:          \t"),
		),
		container.NewVBox(
			widget.NewLabel(secrets.Config().Files.POST_UPLOAD_STEP),
			widget.NewLabel(secrets.Config().Files.VOD_ARCHIVE_DIR),
			widget.NewLabel(secrets.Config().Upload.PLAYLIST_ID),
		),
	)
	ui_context.SetContent(
		func() {},
		widget.NewLabel("Confirm VOD info"),
		content,
	)
}
