package main

import (
	"auto_upload/src/workflow"
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("OBS Auto Upload")

	hello := widget.NewLabel("Fetching VOD info...")
	go func() {
		vod_info := workflow.FetchVodInfo()
		fmt.Println(vod_info)
		hello.SetText(
			fmt.Sprintf("Confirm VOD info:\n%s\n%s\n%s",
				vod_info.Title,
				vod_info.Date,
				vod_info.Ttv_url,
			),
		)
	}()
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Confirm", func() {
			hello.SetText("Good for you :3")
		}),
		widget.NewButton("Close", func() {
			w.Close()
		}),
	))

	w.ShowAndRun()
}
