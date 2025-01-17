package app

import (
	"github.com/JengaMasterG/PalPad/internal/windows"

	"fyne.io/fyne/v2/app"
)

func Start() {
	//true will be data loaded
	//false will be data lost/first time setup
	Load(false)
}

func Load(data bool) {
	a := app.New()
	if !data {
		start := windows.StartPage(a)
		start.Show()
	} else {
		home := windows.HomePage(a)
		home.Show()
	}
	a.Run()
}
