package player

import (
	"fmt"
	"log"
	"strings"

	"palpad/internal/logging"

	"github.com/JengaMasterG/palwrldcmdsgo"
)

func BanKickPlayer(banKick bool, IPAddress string, password string, steamID string) {
	//false is kick, true is ban
	response := ""
	var err error
	banError := "WARN: Could not ban player:"
	kickError := "WARN: Could not kick player:"

	if !banKick {
		log.Printf("INFO: Kicking %s from server", steamID)
		response, err = palwrldcmdsgo.KickPlayer(IPAddress, password, steamID)
	} else {
		log.Printf("INFO: Banning %s from server", steamID)
		response, err = palwrldcmdsgo.BanPlayer(IPAddress, password, steamID)
	}
	if err != nil && !banKick {
		log.Print(kickError, logging.NoConnectionError)
	} else if err != nil && banKick {
		log.Print(banError, logging.NoConnectionError)
	} else {
		log.Print("INFO: ", response)
	}
}

func DataCleaner(listData []string) []string {
	temp := []string{}
	for i := 0; i < len(listData); i++ {
		data := strings.Split(listData[i], "\n")
		for j := 0; j < len(data); j++ {
			if data[j] != "" {
				temp = append(temp, data[j])
			}
		}
	}

	return temp
}

func InfoFormatter(rawData string, cols int) ([][]string, error) {
	listData := strings.Split(rawData, ",")
	cleanData := DataCleaner(listData)
	rows := len(cleanData) / cols
	dataFormatted := [][]string{{"Game Name", "Player ID", "Steam ID"}}

	if len(cleanData)%cols != 0 {
		return nil, fmt.Errorf("the number of columns (%d) does not divide evenly into the max number of data (%d)", cols, len(cleanData))
	}

	for i := 1; i < rows; i++ {
		dataRow := []string{}

		for j := i * 3; j < (i*3)+3; j++ {
			dataRow = append(dataRow, cleanData[j])
		}
		dataFormatted = append(dataFormatted, dataRow)
	}

	return dataFormatted, nil
}

func ShowPlayers(IPAddress string, password string) ([][]string, error) {
	formattedData := [][]string{}
	response, err := palwrldcmdsgo.ShowPlayers(IPAddress, password)
	if err != nil {
		log.Print("WARN: Could not load player list: connection refused")
	} else {
		cols := 3
		formattedData, err = InfoFormatter(response, cols)
		if err != nil {
			log.Println("WARN:", err)
		}
		log.Print("INFO: Loaded Player List")
	}
	return formattedData, err
}
