package main

import (
	"auto_upload/src/workflow"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("OBS Auto Upload")

	hello := widget.NewLabel("Fetching VOD info...")
	go workflow.FetchVodInfo()
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
		widget.NewButton("Close", func() {
			w.Close()
		}),
	))

	w.ShowAndRun()
}
