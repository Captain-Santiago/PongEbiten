package savegame

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var err error

func init() {
	db = nil
	err = nil
}

func InitSaveFile(filePath string) {
	db, err = sql.Open("sqlite3", fmt.Sprintf("file:%s?_foreign_keys=on", filePath))
	if err != nil {
		log.Fatalln(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(64) NOT NULL)`)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS save (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		points INTEGER DEFAULT 0 NOT NULL,
		date DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(user_id) REFERENCES user(id))`)
	if err != nil {
		log.Fatalln(err)
	}
}

func CloseSaveFile() {
	if db != nil {
		err := db.Close()

		if err != nil {
			log.Fatalf("Failed to close DB: %v", err)
		}

		db = nil // reset the db to nil after closing
	}
}
