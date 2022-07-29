package util

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type UIContext struct {
	Window fyne.Window
}

func (ui_context UIContext) SetContent(nextStep func(), content ...fyne.CanvasObject) {
	container := container.NewVBox()

	for i := range content {
		container.Add(content[i])
	}
	if nextStep != nil {
		container.Add(widget.NewButton("Next", nextStep))
	}
	container.Add(widget.NewButton("Close", func() {
		ui_context.Window.Close()
	}))

	ui_context.Window.SetContent(container)
}

func NewUIContext() UIContext {
	app := app.New()
	window := app.NewWindow("OBS Auto Upload")
	window.Resize(fyne.NewSize(600, 600))

	ui_context := UIContext{
		Window: window,
	}
	ui_context.SetContent(nil, widget.NewLabel("Fetching VOD info..."))
	return ui_context
}
