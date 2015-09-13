package main

import (
	"./platform"
	"fmt"
	"os"
	"net/http"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	webArgument := r.URL.Path[1:]
	fmt.Fprintf(w, platform.RandomGameFor(webArgument))
}

func main() {
	len := len(os.Args)
	// Note that first parameter is not app argument
	if ( len == 1 ) {
		log.Println("Start HTTP server on :8080.")
		http.HandleFunc("/", handler)
		http.ListenAndServe(":8080", nil)
	} else if (len == 2 ) {
		argument := os.Args[1]
		if (argument == "--help") {
			printUsage()
		} else {
			fmt.Println(platform.RandomGameFor(argument))
		}
	} else {
		printUsage()
	}
	log.Println("App is closed now")
}

func printUsage() {
	message := "Usage:\n" +
		"\t'gauntlet' - starts the gauntlet web service.\n" +
		"\t'gauntlet [platform]' - prints a random game for provided platform.\n" +
		"\t'gauntlet --help' - prints this message."
	fmt.Println(message)
}