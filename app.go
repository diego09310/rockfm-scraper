package main

import (
	"fmt"
	"os"
	"time"
)

var Version = "dev"

func main() {
	initDb()

	db, console := getArgs()

	prevTitle := getLastSong().Title

	for {
		song := getCurrentSong()
		if song.Title != prevTitle && song.Title != "" && song.Artist != "" {
			if db {
				saveToDb(song)
			}
			if console {
				printToConsole(song)
			}
			prevTitle = song.Title
		}
		time.Sleep(1 * time.Minute)
	}
}

func getArgs() (bool, bool) {
	var db bool
	var console bool
	for _, arg := range os.Args[1:] {
		if arg == "--db" {
			db = true
		}
		if arg == "--console" {
			console = true
		}
		if arg == "--version" || arg == "-v" {
			fmt.Println("rockfmScraper version " + Version)
			os.Exit(0)
		}
	}

	if !db && !console {
		fmt.Println("Usage: rockfmScraper [--db] [--console]")
		os.Exit(1)
	}

	return db, console
}

func printToConsole(song Song) {
	t := time.Now()
	fmt.Println(t.Format("2006-01-02T15:04:05") + " " + song.Artist + " - " + song.Title)
}
