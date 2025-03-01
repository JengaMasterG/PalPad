package windows

import (
	"runtime"

	"github.com/JengaMasterG/PalPad/internal/controllers/server"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var appVersion = "v0.0.3 " + runtime.GOOS

func About(IPAddress string, password string, w fyne.Window) fyne.Window {
	c := w.Canvas()

	data, err := server.Info(IPAddress, password)
	if err != nil {

	}
	verProgram := widget.NewLabelWithStyle(appVersion, fyne.TextAlignCenter, fyne.TextStyle{Bold: false})
	verServer := widget.NewLabelWithStyle(data, fyne.TextAlignCenter, fyne.TextStyle{Bold: false})
	verServer.Wrapping = fyne.TextWrap(fyne.TextWrapBreak)
	container := container.NewVBox(verProgram, verServer)

	c.SetContent(container)
	w.Resize(fyne.NewSize(300, 100))

	return w
}
