package platform

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

func getLocalFilename(platform Platform) (rvalue string) {
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

func GetFilenameFor(platform Platform) (filename string) {
	filename = "systems/" + getLocalFilename(platform) + ".txt"
	return
}