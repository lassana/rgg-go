package main

import (
	"os"
	"bufio"
	"math/rand"
	"time"
	"log"
)

type Platform int

const (
	AMIGA Platform = iota
	C64
	DOS
	DREAMCAST
	GAMEGEAR
	GBA
	GBC
	MAME
	MASTERSYSTEM
	MEGADRIVE
	N64
	NEOGEO
	NES
	PS1
	SATURN
	SNES
	TG16
)

func isValidPlatform(platform string) bool {
	gamesList := []string {
		"amiga",
		"c64",
		"dos",
		"dreamcast",
		"gamegear",
		"gba",
		"gbc",
		"mame",
		"mastrsystem",
		"megadrive",
		"n64",
		"neogeo",
		"nes",
		"ps1",
		"saturn",
		"snes",
		"tg16",
	}
	for _, b := range gamesList {
		if b == platform {
			return true
		}
	}
	return false
}

func readGames(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if ( err != nil) {
		return nil, err
	}
	defer file.Close()

	var games []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		games = append(games, scanner.Text())
	}
	return games, nil
}

func randomGameFor(platform string) string {
	if ( ! isValidPlatform(platform)) {
		log.Printf("'%s' requested as a platform, drop it", platform)
		return ""
	}
	log.Printf("'%s' requested as a platform, handle it", platform)
	filePath := "platforms/" + platform;
	gameList, err := readGames(filePath)
	if ( err == nil) {
		rand.Seed(time.Now().UTC().UnixNano())
		return gameList[rand.Intn(len(gameList))]
	} else {
		log.Printf("Cannot read games list from '%s'", filePath)
		return ""
	}
}
