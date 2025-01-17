package windows

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/JengaMasterG/PalPad/internal/controllers/player"
	"github.com/JengaMasterG/PalPad/internal/controllers/server"
)

var IPAddress, password = "192.168.50.2:25575", "R3dston3$"

func HomePage(a fyne.App) fyne.Window {
	w := a.NewWindow("PalPad")
	w2 := a.NewWindow("About")
	c := w.Canvas()

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
	banBtn := widget.NewButton("Ban", func() { player.BanKickPlayer(true, IPAddress, password, playerId.Text) })
	kickBtn := widget.NewButton("Kick", func() { player.BanKickPlayer(false, IPAddress, password, playerId.Text) })
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
		OnSubmit:   func() { server.Broadcast(IPAddress, password, msg.Text) },
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
	data, err := player.ShowPlayers(IPAddress, password)
	playerList := LoadTable(data)
	//Initial Load of the Table
	if err != nil {
		playerListErr.Show()
		playerList.Hide()
	} else {
		playerListErr.Hide()
		playerList.Show()
	}
	refreshBtn := widget.NewButton("Refresh", func() {
		data, err := player.ShowPlayers(IPAddress, password)
		fmt.Print(data)
		if err != nil {
			playerListErr.Show()
			playerList.Hide()
		} else {
			UpdateTable(playerList, data)
			playerListErr.Hide()
			playerList.Show()
			playerListErr.Refresh()
			playerList.Refresh()
		}
	})
	playerListLayout := container.NewStack(playerList, playerListErr)
	titleLayout := container.NewHBox(title5, layout.NewSpacer(), refreshBtn)
	playerLayout := container.NewGridWithRows(2, titleLayout, playerListLayout)

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
				server.Shutdown(true, IPAddress, password, time.Text, msg2.Text)
			} else {
				time.Validate()
				server.Shutdown(false, IPAddress, password, time.Text, msg2.Text)
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
	leftContainer := container.NewHBox(container.NewVBox(shutdownContainer, broadcastCanvas))
	rightContainer := container.NewGridWithRows(3, title2, playerLayout, banKickContainer)
	content := container.NewGridWithColumns(3, leftContainer, layout.NewSpacer(), rightContainer)

	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File"),

		fyne.NewMenu("Info",
			fyne.NewMenuItem("About", func() { About(w2).Show() })),
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
