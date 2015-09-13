package main

import (
	"./platform"
	"fmt"
	"os"
)

func main() {
	len := len(os.Args)
	if ( len == 1 ) {
		fmt.Println("Web server is not implemented yet.")
	} else if (len == 2 ) { // Note that first parameter is not app argument
		argument := os.Args[1]
		if (argument == "--help") {
			printUsage()
		} else {
			fmt.Println(platform.RandomGameForPlatform(argument))
		}
	} else {
		printUsage()
	}
}

func printUsage() {
	message := "Usage:\n" +
		"\t'gauntlet' - starts the gauntlet web service.\n" +
		"\t'gauntlet [platform]' - prints a random game for provided platform.\n" +
		"\t'gauntlet --help' - prints this message."
	fmt.Println(message)
}