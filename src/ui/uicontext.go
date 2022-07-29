package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type UIContext struct {
	Window fyne.Window
}

func (context UIContext) setContent(nextStep func(), content ...fyne.CanvasObject) {
	container := container.NewVBox()

	for i := range content {
		container.Add(content[i])
	}
	if nextStep != nil {
		container.Add(widget.NewButton("Confirm", nextStep))
	}
	container.Add(widget.NewButton("Close", func() {
		context.Window.Close()
	}))

	context.Window.SetContent(container)
}

func New() UIContext {
	app := app.New()
	window := app.NewWindow("OBS Auto Upload")
	window.Resize(fyne.NewSize(600, 600))

	context := UIContext{
		Window: window,
	}
	context.setContent(nil, widget.NewLabel("Fetching VOD info..."))
	return context
}
