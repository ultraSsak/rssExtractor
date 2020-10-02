package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// executes rssadd (external app!) with the link (adds to newsboat)
func execLinkAppender(link string) {
	_, err := exec.Command("rssadd", link).Output()
	if err != nil {
		log.Fatalln("Can't exec rssadd: ", err)
	}
}

// gets user input
func askForUserInput() int {
	fmt.Print("Which address to add: ")
	var input string
	fmt.Scanln(&input)
	atoi, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalln("User cancel/not number: ", err)
	}
	return atoi
}

// presents link to user
func displayLinks(links []string) {
	for i, link := range links{
		fmt.Printf("[% 3d ] %s\n", i,link)
	}
}

// filter finds links that might be atom/rss ones, replaces input in-place
func filter(links []string, filterValues ...string) []string {
	i:=0
	for _, s := range links {
		for _, f := range filterValues {
			if strings.Contains(s, f) {
				links[i] = s
				i++
				break
			}
		}
	}
	return links[:i]
}

// findLinks appends to links each link found in n and returns the result.
func findLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "link") {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = findLinks(links, c)
	}
	return links
}
// grabs and parses the web page
func grabThePage() *html.Node {
	// grab url from args
	var link = os.Args[1]

	// add https in front if necessary
	if !strings.HasPrefix(link,"http") {
		link = "https://" + link
	}
	// make request
	get, err := http.Get(link)
	if err != nil {
		log.Fatalf("can't load page: %s\n", err)
	}
	// grab the response body
	defer get.Body.Close()
	node, err := html.Parse(get.Body)

	if err != nil {
		log.Fatalf("Cant parse html: %s", err)
	}
	return node
}
