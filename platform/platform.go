package platform

import (
	"os"
	"bufio"
	"fmt"
	"math/rand"
	"time"
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

func localFilename(platform Platform) (rvalue string) {
	switch platform {
	case AMIGA:
		rvalue = "amiga"
	case C64:
		rvalue = "c64"
	case DOS:
		rvalue = "dos"
	case DREAMCAST:
		rvalue = "dreamcast"
	case GAMEGEAR:
		rvalue = "gamegear"
	case GBA:
		rvalue = "gba"
	case GBC:
		rvalue = "gbc"
	case MAME:
		rvalue = "mame"
	case MASTERSYSTEM:
		rvalue = "mastrsystem"
	case MEGADRIVE:
		rvalue = "megadrive"
	case N64:
		rvalue = "n64"
	case NEOGEO:
		rvalue = "neogeo"
	case NES:
		rvalue = "nes"
	case PS1:
		rvalue = "ps1"
	case SATURN:
		rvalue = "saturn"
	case SNES:
		rvalue = "snes"
	case TG16:
		rvalue = "tg16"
	default:
		rvalue = ""
	}
	return
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

func RandomGameForPlatform(platform string) string {
	filePath := "retro-gauntlet/cgi-bin/systems/" + platform;
	gameList, err := readGames(filePath)
	if ( err == nil) {
		rand.Seed(time.Now().UTC().UnixNano())
		return gameList[rand.Intn(len(gameList))]
	} else {
		fmt.Printf("Cannot read games list from '%s'", filePath)
		return ""
	}
}

func RandomGameFor(platform Platform) string {
	return RandomGameForPlatform(localFilename(platform))
}