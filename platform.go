package main

import (
	"os"
	"bufio"
	"math/rand"
	"time"
	"log"
	"sync"
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

var lock sync.Mutex

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
		"mastersystem",
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
	} else {

	}
	lock.Lock()
	log.Printf("'%s' requested as a platform, handle it", platform)
	filePath := "platforms/" + platform;
	gameList, err := readGames(filePath)
	var rvalue string
	if ( err == nil) {
		rand.Seed(time.Now().UTC().UnixNano())
		rvalue = gameList[rand.Intn(len(gameList))]
	} else {
		log.Printf("Cannot read games list from '%s'", filePath)
		rvalue = ""
	}
	lock.Unlock()
	return rvalue
}
