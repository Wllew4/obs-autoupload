package w_vod

import (
	"auto_upload/src/util"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func UI_vod(ui_context util.UIContext, vod_info VOD, nextStep func()) {
	content := container.New(layout.NewFormLayout(),
		container.NewVBox(
			widget.NewLabel("Title:\t"),
			widget.NewLabel("Date: \t"),
			widget.NewLabel("Path: \t"),
			widget.NewLabel("Url:  \t"),
		),
		container.NewVBox(
			util.PopulatedEntry(&vod_info.Title),
			util.PopulatedEntry(&vod_info.Date),
			util.PopulatedEntry(&vod_info.Path),
			util.PopulatedEntry(&vod_info.Ttv_url),
		),
	)
	ui_context.SetContent(
		func() { go nextStep() },
		widget.NewLabel("Confirm VOD info"),
		content,
	)
}
