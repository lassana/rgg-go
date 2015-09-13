package main

import (
	"./platform"
	"fmt"
	"os"
	"net/http"
	"log"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	webArgument := r.URL.Path[1:]
	//webArgument := r.URL.Path[strings.LastIndex('/', r.URL)]
	fmt.Fprintf(w, platform.RandomGameFor(strings.Trim(webArgument, "platform/")))
}

func main() {
	len := len(os.Args)
	// Note that first parameter is not app argument
	if ( len == 1 ) {
		log.Println("Start HTTP server on :8080.")
		http.HandleFunc("/platform/", handler)
		http.Handle("/", http.FileServer(http.Dir("static/")))
		panic(http.ListenAndServe(":8080", nil))
		//panic(http.ListenAndServe(":8080", http.FileServer(http.Dir("static/"))))
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