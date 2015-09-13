package main

import (
	"fmt"
	"os"
	"net/http"
	"log"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	webArgument := r.URL.Path[1:]
	log.Printf("Handle request '%s'", webArgument)
	webArgument = strings.Replace(webArgument, "platform/", "", 1)
	//webArgument = strings.Replace(webArgument, "/", "", 1)
	fmt.Fprintf(w, randomGameFor(webArgument))
}

func main() {
	len := len(os.Args)
	// Note that first parameter is not app argument
	if ( len == 1 ) {
		log.Println("Start HTTP server...")
		http.HandleFunc("/platform/", handler)
		http.Handle("/", http.FileServer(http.Dir("static/")))
		bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))
		fmt.Printf("listening on %s...", bind)
		panic(http.ListenAndServe(bind, nil))
		//panic(http.ListenAndServe(":8080", http.FileServer(http.Dir("static/"))))
	} else if (len == 2 ) {
		argument := os.Args[1]
		if (argument == "--help") {
			printUsage()
		} else {
			fmt.Println(randomGameFor(argument))
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