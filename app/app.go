package app

import (
	"log"

	"github.com/JengaMasterG/PalPad/internal/controllers/database"
	"github.com/JengaMasterG/PalPad/internal/windows"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func Init() {
	a := app.New()
	init := windows.InitPage("Booting PalPad...", a)
	init.Show()
	go Start(a, init)
	a.Run()
	a.Quit()
	}

func Start(a fyne.App, init fyne.Window) {
	
	// Check if chai database was created previously
	data, err := database.GetData(0)
	if err != nil {
		database.InitDB()
	}

	//Check if chai database has user data in it
	data, err = database.GetData(1)
	if err != nil {
		log.Print("INFO: Start() ", err.Error())
	}

	//0 will be data loaded
	//1 will be data lost/first time setup
	if data.IPAddress == "" {
		Load(data.ID, 1, a, init)
	} else {
		Load(data.ID, 0, a, init)
	}
}

func Load(id int, dataFlag int, a fyne.App, init fyne.Window) {
	switch dataFlag {
	case 0:
		home := windows.HomePage(id, a)
		home.Show()
		init.Close()
	case 1:
		start := windows.StartPage(a)
		start.Show()
		init.Close()
	default:
		start := windows.StartPage(a)
		start.Show()
		init.Close()
	}
	//box.RemoveAll() uncomment if debugging
}
