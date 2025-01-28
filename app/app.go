package app

import (
	"log"

	"github.com/JengaMasterG/PalPad/internal/controllers/database"
	"github.com/JengaMasterG/PalPad/internal/windows"

	"fyne.io/fyne/v2/app"
)

func Start() {
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
		Load(data.ID, 1)
	} else {
		Load(data.ID, 0)
	}
}

func Load(id int, dataFlag int) {
	a := app.New()
	switch dataFlag {
	case 0:
		home := windows.HomePage(id, a)
		home.Show()
	case 1:
		start := windows.StartPage(a)
		start.Show()
	default:
		start := windows.StartPage(a)
		start.Show()
	}
	a.Run()
	//box.RemoveAll() uncomment if debugging
}
