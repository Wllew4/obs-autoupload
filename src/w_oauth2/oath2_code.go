package w_oauth2

import (
	"auto_upload/src/util"
	"net/http"

	"fyne.io/fyne/v2/widget"
	"golang.org/x/oauth2"
)

func receiveCode(
	ui_context util.UIContext,
	oauth2CodeCallback func(
		string,
		string,
		*oauth2.Config,
		func(*http.Client)),
	cacheFile string,
	config *oauth2.Config,
	nextStep func(*http.Client),
) {
	entry := widget.NewEntry()
	ui_context.SetContent(
		func() {
			oauth2CodeCallback(
				entry.Text,
				cacheFile,
				config,
				nextStep,
			)
		},
		widget.NewLabel("Authorize app and enter code:"),
		entry,
	)
}
