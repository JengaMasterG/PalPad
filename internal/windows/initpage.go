package windows

import (
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func InitPage(initStatus string, a fyne.App) fyne.Window {
	w := a.NewWindow("PalPad")
	c := w.Canvas()
	appVersion := "v0.0.5 " + runtime.GOOS

	verProgram := widget.NewLabelWithStyle(appVersion, fyne.TextAlignCenter, fyne.TextStyle{Bold: false})
	initRow := widget.NewLabelWithStyle("Starting up...", fyne.TextAlignCenter, fyne.TextStyle{Bold: false})
	initRow.Wrapping = fyne.TextWrap(fyne.TextWrapBreak)
	container := container.NewVBox(verProgram, initRow)

	c.SetContent(container)
	w.Resize(fyne.NewSize(300, 100))

	return w
}
