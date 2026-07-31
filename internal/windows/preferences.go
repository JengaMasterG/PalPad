package windows

import (
	"log"
	"strconv"

	"github.com/JengaMasterG/PalPad/internal/controllers/database"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Preferences(w fyne.Window) fyne.Window {
	c := w.Canvas()
	//db := database.GetDB()

	/**
	-1 is used to tell the GetData function to grab all the data from the table
	**/
	data, err := database.GetData(-1)
	if err != nil {
		log.Print("WARN: ", err)
	}else{
		log.Print("INFO: ", data)
	}
	loginLbl := widget.NewLabel("Server Settings")
	//loginLbl.Alignment = fyne.TextAlignCenter

	
	ipAddr := widget.NewEntry()
	pass := widget.NewPasswordEntry()
	ipAddr.SetPlaceHolder("")
	pass.SetPlaceHolder("")
	issue := widget.NewLabel("")
	issue.Alignment = fyne.TextAlignCenter
	issue.Wrapping = fyne.TextWrap(fyne.TextWrapBreak)
	idSelect := widget.NewSelect([]string{strconv.Itoa(data.ID)}, func(value string) {
		//Convert the selected ID back to an int
		selectedId, err := strconv.Atoi(value)
		if err != nil{
			log.Print("WARN: ", err)
			issue.SetText("Error!")
		}else{
			data, err := database.GetData(selectedId)
			if err != nil {
				log.Print("WARN: ", err)
			}else{
				ipAddr.SetText(data.IPAddress)
				pass.SetText(data.AdminPassword)
			}
		}
	})

	updateBtn := widget.NewButton("Update", func() {
		id, err :=strconv.Atoi(idSelect.Selected)
		if err != nil{
			log.Print("WARN: ", err)
		}else{
			err := database.UpdateData(id, ipAddr.Text, pass.Text)
			log.Print("From Input ", ipAddr.Text, " ", pass.Text)
				if err != nil {
				issue.SetText(err.Error())
			} else {
				data, err = database.GetData(id)
				if err != nil{
					log.Print("WARN: ", err)
				}else{
					ipAddr.SetText(data.IPAddress)
					pass.SetText(data.AdminPassword)
					issue.SetText("Server Info Updated!")
				}
			}
		}
	})
	clearBtn := widget.NewButton("Cancel", func() {
		ipAddr.SetText("")
		pass.SetText("")
		issue.SetText("")
	})
	btnLayout := container.NewGridWithColumns(2, clearBtn, updateBtn)
	component := container.NewGridWithRows(6, loginLbl, idSelect, ipAddr, pass, btnLayout, issue)
	container := container.NewGridWithColumns(3, layout.NewSpacer(), component, layout.NewSpacer())

	c.SetContent(container)
	w.Resize(fyne.NewSize(600, 100))

	return w
}
