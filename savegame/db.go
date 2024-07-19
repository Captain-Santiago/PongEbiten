package savegame

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
}

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

// func GetIdByName(name string) (int, error) {
// 	rows, err := db.Query(`SELECT id FROM user WHERE name = ? LIMIT = 1`, name)
// 	if err != nil {
// 		return -1, err
// 	}

// 	var user_id []int
// 	err = rows.Scan(&user_id)
// 	if err != nil {
// 		return -1, err
// 	}

// 	user_not_found := &ErrUserNotFound{name: name}
// 	if len(user_id) == 0 {
// 		return 0, user_not_found
// 	}

// 	return user_id[0], nil
// }

// func InsertNewElement(save Save, user_id int) error {
// 	// Create user if not exists
// 	user_id, err := GetIdByName(save.Name)
// 	if errors. {

// 	}

// 	// Insert new data into the user
// 	query, err := db.Query(`INSERT INTO save (user_id, points)
// 							VALUES(?,?);`, user_id, save.Points)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func CloseSaveFile() {
	if db != nil {
		err := db.Close()

		if err != nil {
			log.Fatalf("Failed to close DB: %v", err)
		}

		db = nil // reset the db to nil after closing
	}
}
