package main

import (
	"auto_upload/src/workflow"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func populatedEntry(text *string) *widget.Entry {
	o := widget.NewEntry()
	o.Bind(binding.BindString(text))
	return o
}

func main() {
	a := app.New()
	w := a.NewWindow("OBS Auto Upload")
	w.Resize(fyne.NewSize(600, 600))

	hello := widget.NewLabel("Fetching VOD info...")

	content := container.NewVBox(
		hello,
		widget.NewButton("Confirm", func() {
			hello.SetText("Good for you :3")
		}),
		widget.NewButton("Close", func() {
			w.Close()
		}),
	)

	go func() {
		vod_info := workflow.FetchVodInfo()
		fmt.Println(vod_info)
		hello.SetText("Confirm VOD info:")
		a := container.NewHSplit(
			container.NewVBox(
				widget.NewLabel("Title:\t"),
				widget.NewLabel("Date: \t"),
				widget.NewLabel("Path: \t"),
				widget.NewLabel("Url:  \t"),
			),
			container.NewVBox(
				populatedEntry(&vod_info.Title),
				populatedEntry(&vod_info.Date),
				populatedEntry(&vod_info.Path),
				populatedEntry(&vod_info.Ttv_url),
			),
		)
		a.SetOffset(0.1)
		content.Add(a)
	}()

	w.SetContent(content)
	w.ShowAndRun()
}
