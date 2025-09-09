package main

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "./data/"
const dbFile = "./data/songs.db"

func initDb() {
	if _, err := os.Stat(dbFile); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(dbPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
		os.OpenFile(dbFile, os.O_RDONLY, 0666)
	}

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS songs (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, artist TEXT, timestamp DATETIME DEFAULT CURRENT_TIMESTAMP);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}

func saveToDb(song Song) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	stmt, err := tx.Prepare("INSERT INTO songs(title, artist) VALUES(?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(song.Title, song.Artist)
	if err != nil {
		panic(err)
	}
	tx.Commit()
}
