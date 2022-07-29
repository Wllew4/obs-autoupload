package ui

import (
	"fyne.io/fyne/v2/widget"
	"golang.org/x/oauth2"
)

func ReceiveCode(
	ui_context UIContext,
	tokenCallback func(*oauth2.Config, string, string) *oauth2.Token,
	config *oauth2.Config,
	cacheFile string,
) {

	entry := widget.NewEntry()
	ui_context.SetContent(func() {
		tokenCallback(config, entry.Text, cacheFile)
	}, entry)
}
