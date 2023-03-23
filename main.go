package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/methosi/fsmk_test/playlist"
)

var noOfSong int

func init() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide number of song")
		os.Exit(1)
	}

	var err error
	noOfSong, err = strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if noOfSong > 500 {
		fmt.Println("Number of song should be less than 500")
		os.Exit(1)
	}
}

func main() {
	var (
		reader = bufio.NewReader(os.Stdin)
	)

	// Init playlist
	playlist := playlist.NewPlaylist(noOfSong)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		// Trim the input
		input = strings.TrimSpace(input)

		// Check if input is empty
		if input == "" {
			continue
		}

		// Split the input
		words := strings.Split(input, " ")
		if len(words) < 2 {
			fmt.Println("Invalid input, please input song title")
			continue
		}

		switch strings.ToUpper(words[0]) {
		case "PLAY":
			// Get the song title
			songTitle := words[1]
			playlist.PlaySong(songTitle)
			fmt.Println(songTitle + " Played")
			fmt.Println("Playlist: ", playlist.GetSongList())
		default:
			fmt.Println("Invalid input, please input PLAY")
		}
	}
}
