package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDatabase() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "golangTranslate",
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingerr := db.Ping()
	if pingerr != nil {
		log.Fatal(pingerr)
	}
	fmt.Println("Connected")
	word, err := getWord(RU_LOCALE, "абракадабра")
	fmt.Println(word, " ", err)
	word, err = getWord(EN_LOCALE, "hello")
	fmt.Println(word, " ", err)
}

//Returns a word translation from database
func getWord(locale, input string) (word DBWord, err error) {
	var condition string
	switch locale {
	case RU_LOCALE:
		condition = "ru"
	case EN_LOCALE:
		condition = "en"
	default:
		{
			err = fmt.Errorf("Incorrect locale (recieved {%s})", locale)
			return
		}
	}
	queryText := fmt.Sprintf("SELECT * FROM dictionary WHERE %s=?", condition)
	var row *sql.Row
	row = db.QueryRow(queryText, input)
	if err = row.Scan(&word.ID, &word.EnTranslation, &word.RuTranslation, &word.AppealCounter); err != nil {
		return
	}
	incrementCount(int(word.ID))
	return
}

func incrementCount(id int) (err error) {
	_, err = db.Exec("UPDATE dictionary SET appeal_count = appeal_count + 1 WHERE id = ?", id)
	return
}

func addWord(ruTranslation, enTranslation string) error {
	_, err := db.Exec("INSERT INTO dictionary (en, ru, appeal_count) VALUES (?, ?, ?)", enTranslation, ruTranslation, 1)
	if err != nil {
		return err
	}
	return nil

}
