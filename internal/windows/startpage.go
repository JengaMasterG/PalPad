package windows

import (
	"log"

	"github.com/JengaMasterG/PalPad/internal/controllers/database"
	"github.com/JengaMasterG/PalPad/internal/controllers/server"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func StartPage(a fyne.App) fyne.Window {
	w := a.NewWindow("PalPad")
	c := w.Canvas()
	db := database.GetDB()

	loginLbl := widget.NewLabel("Login")
	loginLbl.Alignment = fyne.TextAlignCenter

	ipAddr := widget.NewEntry()
	pass := widget.NewPasswordEntry()
	ipAddr.SetPlaceHolder("IP Address:Port")
	pass.SetPlaceHolder("Enter Password")

	issue := widget.NewLabel("")
	issue.Alignment = fyne.TextAlignCenter
	issue.Wrapping = fyne.TextWrap(fyne.TextWrapBreak)

	connectBtn := widget.NewButton("Connect", func() {
		data, err := server.Info(ipAddr.Text, pass.Text)
		log.Print("From Input", ipAddr.Text, pass.Text)
		if err != nil {
			issue.SetText(err.Error())
		} else {
			log.Print("INFO: ", data)
			err := database.SetData(db, ipAddr.Text, pass.Text)
			if err != nil {
				issue.SetText(err.Error())
			} else {
				db.Close()
				w.Hide()
				home := HomePage(1, a)
				home.Show()
			}
		}
	})
	clearBtn := widget.NewButton("Clear", func() {
		ipAddr.SetText("")
		pass.SetText("")
		issue.SetText("")
	})
	btnLayout := container.NewGridWithColumns(2, clearBtn, connectBtn)
	component := container.NewGridWithRows(5, loginLbl, ipAddr, pass, btnLayout, issue)
	container := container.NewGridWithColumns(3, layout.NewSpacer(), component, layout.NewSpacer())

	c.SetContent(container)

	return w
}
