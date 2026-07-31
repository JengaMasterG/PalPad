package database

import (
	"log"
	"strconv"

	"github.com/chaisql/chai"
)

type ServerInfo struct {
	ID            int
	IPAddress     string
	AdminPassword string
}

func InitDB() {
	db := GetDB()
	defer db.Close()

	err := db.Exec(`
	CREATE TABLE serverInfo (
		ID				INT		PRIMARY KEY,
		IPAddress		TEXT	NOT NULL UNIQUE,
		AdminPassword	TEXT	NOT NULL UNIQUE,
	
		CHECK (len(IPAddress) > 0),
		CHECK (len(AdminPassword) > 0)
	)
	`)

	if err != nil {
		log.Print("WARN: ", err)
	}
}

func GetDB() *chai.DB {
	db, err := chai.Open("data")
	if err != nil {
		log.Print("INFO: Could not find database; ", err.Error())
	}

	return db
}

func GetData(id int) (ServerInfo, error) {
	/**
	Logic
	If id is greater than or equal to 0, pull the data just for that id
	If id is less than 0, pull all the data from the table
	**/
	
	db := GetDB()
	defer db.Close()
	var query string

	if id >= 0{
		query = "SELECT ID, IPAddress, AdminPassword FROM serverInfo WHERE ID = " + strconv.Itoa(id)
	} else {
		query = "SELECT ID, IPAddress, AdminPassword FROM serverInfo"
	}
	var s ServerInfo

	stream, err := db.Query(query)
	if err != nil {
		log.Print("INFO: GetDB:", err.Error())
		return s, err
	} else {
		log.Print("INFO: GetDB: Query Sent")
	}
	defer stream.Close()

	err = stream.Iterate(func(r *chai.Row) error {
		err = r.Scan(&s.ID, &s.IPAddress, &s.AdminPassword) // Had to manually set values into struct
		//err = r.StructScan(&s) <--Doesn't read into a struct format
		if err != nil {
			log.Print("WARN: GetDB:", err.Error())
		} else {
			log.Print("INFO: GetDB: Database loaded")
		}
		return nil
	})

	return s, err
}

func SetData(db *chai.DB, IPAddr string, Password string) error {
	//Get the number of IDs already in the database
	var id int
	ipAddr := IPAddr
	pass := Password

	rows, err := db.Query("SELECT ID FROM serverInfo")
	if err != nil {
		log.Print("INFO: ", err.Error())
	}
	defer rows.Close()

	err = rows.Iterate(func(r *chai.Row) error {
		err = r.Scan(&id)
		return nil
	})

	//for now, grab the last id and add 1 for new ID
	log.Print("INFO: ServerInfo ID: ", id)

	newId := id + 1

	s := ServerInfo{
		ID:            newId,
		IPAddress:     ipAddr,
		AdminPassword: pass,
	}

	log.Print("INFO: New data to be added: ", s)

	err = db.Exec(`INSERT INTO serverInfo (ID, IPAddress, AdminPassword) VALUES (?, ?, ?)`, s.ID, s.IPAddress, s.AdminPassword)
	if err != nil {
		log.Print("INFO: ", err.Error())
	}
	return err
}

func UpdateData(id int, ipAddr string, Password string) (error){
	db := GetDB()
	defer db.Close()

	err := db.Exec(`UPDATE serverInfo SET IPAddress = ?, AdminPassword = ? WHERE ID = ?`, ipAddr, Password, id)
	if err != nil {
		log.Print("INFO: UpdateDB:", err.Error())
		return err
	} else {
		log.Print("INFO: UpdateDB: Updated server info for ID: ", strconv.Itoa(id))
		return err
	}	
}
