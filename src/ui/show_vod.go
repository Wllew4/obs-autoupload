package ui

import (
	"auto_upload/src/util"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func populatedEntry(text *string) *widget.Entry {
	o := widget.NewEntry()
	o.Bind(binding.BindString(text))
	return o
}

func ShowVOD(ui_context UIContext, vod_info util.VOD, nextStep func()) {
	content := container.New(layout.NewFormLayout(),
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
	ui_context.SetContent(
		func() { go nextStep() },
		widget.NewLabel("Confirm VOD info"),
		content,
	)
}
