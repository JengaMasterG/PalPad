package server

import (
	"log"
	"strconv"
	"time"

	"github.com/JengaMasterG/PalPad/internal/logging"

	"github.com/JengaMasterG/palwrldcmdsgo"
)

func Broadcast(IPAddress string, password string, message string) {
	err := palwrldcmdsgo.Broadcast(IPAddress, password, message)
	if err != nil {
		log.Print("WARN: ", err)
	}
	log.Print("INFO: Message sent\n", "INFO: ", message)
}

func Info(IPAddress string, password string) (string, error) {
	response, err := palwrldcmdsgo.Info(IPAddress, password)
	if err != nil {
		log.Print("WARN: ", logging.NoConnectionError)
		response = "Could not access server information"
	}
	return response, err
}

func Save(IPAddress string, password string) (string, error) {
	response, err := palwrldcmdsgo.Save(IPAddress, password)
	if err != nil {
		log.Print("FATAL: Could not save data; ", logging.NoConnectionError)
	} else {
		log.Print("INFO: ", response)
	}
	return response, err
}

func Shutdown(force bool, IPAddress string, password string, seconds string, message string) {
	//TODO: rework to cancel a shutdown (Time before shutdown)
	//TODO: If Save throws error, stop shutdown
	//false is safe shutdown, true is force shutdown
	minTimer := 10
	response := ""
	sec, err := strconv.Atoi(seconds)
	shutdownWarn := "WARN: Could not shutdown server: "
	if err != nil {
		log.Print(shutdownWarn, "Invalid timer format")
	} else {
		timer := 0
		if minTimer >= timer {
			timer = 0
		} else {
			timer = sec - minTimer
		}
		if !force {
			response, err = palwrldcmdsgo.Shutdown(IPAddress, password, seconds, message)
			if err != nil {
				log.Print(shutdownWarn, logging.NoConnectionError)
			} else {
				log.Print(response)
				saveTimer := time.NewTimer(time.Duration(timer) * time.Second)
				<-saveTimer.C
				Save(IPAddress, password)
			}
		} else {
			log.Print("WARN: FORCE SHUTDOWN ISSUED FROM PALPAD")
			response, err = palwrldcmdsgo.DoExit(IPAddress, password)
			if err != nil {
				log.Print(shutdownWarn, logging.NoConnectionError)
			} else {
				log.Print("WARN: ", response)
			}
		}
	}
}
