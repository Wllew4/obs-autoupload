package ui

import (
	"net/http"

	"fyne.io/fyne/v2/widget"
	"golang.org/x/oauth2"
)

func ReceiveCode(
	ui_context UIContext,
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
