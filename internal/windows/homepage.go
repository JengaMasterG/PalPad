package windows

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/JengaMasterG/PalPad/internal/controllers/database"
	"github.com/JengaMasterG/PalPad/internal/controllers/player"
	"github.com/JengaMasterG/PalPad/internal/controllers/server"
)

func HomePage(id int, a fyne.App) fyne.Window {
	w := a.NewWindow("PalPad")
	w2 := a.NewWindow("About")
	c := w.Canvas()
	data, err := database.GetData(id)
	ipAddrData := data.IPAddress
	passData := data.AdminPassword

	//title1 := widget.NewLabel("Server Info")
	title2 := widget.NewLabel("Player Management")
	title3 := widget.NewLabel("Shutdown Server")
	title4 := widget.NewLabel("Send a message to the PalWorld Server")
	title5 := widget.NewLabel("Player List")

	//need to create form widgets for each command that needs and entry
	//Ban and Kick Players
	//Vertical Box conatiner
	//|------PlayerID Field------|
	//|BanBtn|			 |KickBtn| <-- GridLayout w Btns wrapped in HBox
	playerId := widget.NewEntry()
	playerId.SetPlaceHolder("User's Steam ID")
	banBtn := widget.NewButton("Ban", func() { player.BanKickPlayer(true, ipAddrData, passData, playerId.Text) })
	kickBtn := widget.NewButton("Kick", func() { player.BanKickPlayer(false, ipAddrData, passData, playerId.Text) })
	btnRow := container.New(layout.NewHBoxLayout(), kickBtn, banBtn)
	banKickContainer := container.NewVBox(playerId, btnRow)

	//Broadcast Container
	//|---Message---|||SendBtn| <--Inverted form layout
	msg := widget.NewEntry()
	msg.SetPlaceHolder("Enter text...")
	msgForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "SYSTEM Message", Widget: msg},
		},
		OnSubmit:   func() { server.Broadcast(ipAddrData, passData, msg.Text) },
		SubmitText: "Send Msg",
		OnCancel:   func() { msg.SetText("") },
		CancelText: "Clear",
	}
	//btnLayout := container.NewGridWithColumns(2, layout.NewSpacer(), sendBtn)
	msglayout := container.NewVBox(msgForm)
	broadcastCanvas := container.NewVBox(title4, msglayout)

	//List Players
	//|--Label--| |RefreshBtn|
	//|--------Table---------|
	//Cols needs to divide evely into data
	playerListErr := widget.NewLabel("Unable to Load Player List")
	tableData, err := player.ShowPlayers(ipAddrData, passData)
	playerList := LoadTable(tableData)
	//Initial Load of the Table
	if err != nil {
		playerListErr.Show()
		playerList.Hide()
	} else {
		playerListErr.Hide()
		playerList.Show()
	}
	refreshBtn := widget.NewButton("Refresh", func() {
		tableData, err := player.ShowPlayers(ipAddrData, passData)
		fmt.Print(tableData)
		if err != nil {
			playerListErr.Show()
			playerList.Hide()
		} else {
			UpdateTable(playerList, tableData)
			playerListErr.Hide()
			playerList.Show()
			playerListErr.Refresh()
			playerList.Refresh()
		}
	})
	playerListLayout := container.NewStack(playerList, playerListErr)
	titleLayout := container.NewVBox((container.NewHBox(title5, layout.NewSpacer(), refreshBtn)))
	playerLayout := container.NewGridWithRows(2, titleLayout, playerListLayout)

	//Save Server
	saveLabel := widget.NewLabel("Save Server State")
	saveStatus := widget.NewLabel("")
	saveBtn := widget.NewButton("Save", func() {
		response, err := server.Save(ipAddrData, passData)
		if err != nil {
			saveStatus.SetText("Could not save data!")
		} else {
			saveStatus.SetText(response)
		}
		saveStatus.Refresh()
	})
	saveLayout := container.NewHBox(saveLabel, layout.NewSpacer(), saveBtn)
	saveContainer := container.NewGridWithRows(2, saveLayout, saveStatus)

	//Shutdown Server
	time := widget.NewEntry()
	time.Validator = func(str string) error {
		if _, err := strconv.Atoi(str); err != nil {
			return fmt.Errorf("Must be numbers only")
		}
		return nil
	}
	msg2 := widget.NewEntry()
	msg2.SetPlaceHolder("Enter text...")

	forceShutdown := widget.NewCheck("", func(value bool) {
		if value {
			time.SetText("0")
			msg2.SetText("")
			time.Disable()
			msg2.Disable()
		} else {
			time.Enable()
			msg2.Enable()
		}
	})

	shutdown := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Time (in seconds)", Widget: time},
			{Text: "SYSTEM Message", Widget: msg2},
			{Text: "Force Shutdown?", Widget: forceShutdown},
		},
		OnSubmit: func() {
			if forceShutdown.Checked {
				server.Shutdown(true, ipAddrData, passData, time.Text, msg2.Text)
			} else {
				time.Validate()
				server.Shutdown(false, ipAddrData, passData, time.Text, msg2.Text)
			}
		},
		SubmitText: "Shutdown",
		OnCancel: func() {
			time.SetText("")
			msg2.SetText("")
			forceShutdown.SetChecked(false)
		},
		CancelText: "Clear",
	}
	shutdownLayout := container.NewVBox(shutdown)
	shutdownContainer := container.NewVBox(title3, shutdownLayout)

	//Load the content
	leftContainer := container.NewVBox(saveContainer, shutdownContainer, broadcastCanvas)
	rightContainer := container.NewGridWithRows(3, title2, playerLayout, banKickContainer)
	content := container.NewGridWithColumns(3, leftContainer, layout.NewSpacer(), rightContainer)

	//TODO: Add Reset data function
	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File"),

		fyne.NewMenu("Info",
			fyne.NewMenuItem("About", func() { About(ipAddrData, passData, w2).Show() })),
	))

	c.SetContent(content)

	return w
}

func LoadTable(data [][]string) *widget.Table {
	cols := 3
	rows := len(data)
	table := widget.NewTable(func() (int, int) { return rows, cols },
		func() fyne.CanvasObject {
			return widget.NewLabel(strings.Repeat("", 60))
		},
		func(id widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[id.Row][id.Col])
		},
	)

	return table
}

func UpdateTable(table *widget.Table, data [][]string) {
	//Get table length to compare with the lenght of new data
	row, col := table.Length()
	rowData := len(data)

	//adjust the number of rows of the table to match the new data
	if rowData > row || rowData < row {
		table.Length = func() (int, int) { return rowData, col }
	}

	//after the table is the right size, we populate it with the new data
	table.UpdateCell = func(id widget.TableCellID, o fyne.CanvasObject) {
		defer func() {
			if err := recover(); err != nil {
				o.(*widget.Label).SetText("")
			}
		}() //handler in case of a panic
		o.(*widget.Label).SetText(data[id.Row][id.Col])
	}
}
