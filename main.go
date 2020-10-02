// finds and extracts link's for atom/rss readers
package main

import (
	"log"
	"os"
)

// entry point
func main() {
	// guard for usage, exits when wrong
	usage()

	// fetch the web page
	doc := grabThePage()

	// find links, then filter them by patterns
	links := findLinks(nil, doc)
	links = filter(links, "rss", "atom", "feed")
	displayLinks(links)

	// get user input and execute adding address
	linkNum := askForUserInput()
	execLinkAppender(links[linkNum])
}

// usage prints syntax
func usage() {
	if len(os.Args) < 2 {
		log.SetFlags(0)
		log.Fatalf("USAGE: %s https://link.to.page.com\n", os.Args[0])
	}
}
