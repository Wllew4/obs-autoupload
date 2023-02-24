package w_upload

import (
	"auto_upload/src/util"
	"net/url"

	"fyne.io/fyne/v2/widget"
)

func ui_uploading(ui_context util.UIContext) {
	ui_context.SetContent(
		nil,
		widget.NewLabel("Uploading...\nThis may take awhile..."),
	)
}

func ui_uploaded(ui_context util.UIContext, nextStep func(string), id string) {
	url, err := url.Parse("https://youtube.com/watch?v=" + id)
	util.CheckErr(err)
	ui_context.SetContent(
		func() { nextStep(id) },
		widget.NewLabel("Success!"),
		widget.NewHyperlink("https://youtube.com/watch?v="+id, url),
	)
}
