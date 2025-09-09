package main

import (
	"bufio"
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"gopkg.in/vansante/go-ffprobe.v2"
)

const rockfmUrl = "https://rockfm-cope.flumotion.com/"
const chunksUrl = rockfmUrl + "chunks.m3u8"

func getCurrentSong() Song {
	file, err := http.Get(chunksUrl)
	if err != nil {
		log.Fatalln(err)
		return Song{}
	}
	defer file.Body.Close()
	scanner := bufio.NewScanner(file.Body)
	var audioFile string
	for scanner.Scan() {
		if scanner.Text()[0] != '#' {
			audioFile = scanner.Text()
			break
		}
	}
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()
	data, err := ffprobe.ProbeURL(ctx, rockfmUrl+audioFile)
	if err != nil {
		log.Fatalln("Error probing URL: ", err)
		return Song{}
	}
	title, err := data.Format.TagList.GetString("title")
	if err != nil {
		log.Fatalln("Error extracting title: ", err)
		return Song{}
	}
	song := strings.Split(strings.TrimSuffix(title, " <nil>"), " - ")
	return Song{Title: song[0], Artist: song[1]}
}
