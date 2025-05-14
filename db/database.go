package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./iot.db")
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS sensor_data (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		power TEXT,
		voltage TEXT,
		current TEXT
	);
	`
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database initialized")
}

func InsertSensor(name, power, voltage, current string) {
	stmt, err := DB.Prepare("INSERT INTO sensor_data(name, power, voltage, current) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = stmt.Exec(name, power, voltage, current)
	if err != nil {
		log.Println(err)
	}
}
