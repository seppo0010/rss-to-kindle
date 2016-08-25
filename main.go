package main

import (
	"fmt"
	"os"
	"rss-to-kindle/content"
	"rss-to-kindle/kindle"

	"github.com/fatih/color"
)

func printConfiguration(email string, links []string) {
	color.New(color.FgCyan).Add(color.Underline).Println("Configuration")
	fmt.Println("Email:\t" + email)

	if len(links) > 0 {
		fmt.Printf("Links:")
		for _, link := range links {
			fmt.Println("\t* " + link)
		}
	}
}

func main() {
	server := os.Args[1]
	port := os.Args[2]
	fromEmail := os.Args[3]
	password := os.Args[4]
	toEmail := os.Args[5]

	links := os.Args[6:]

	printConfiguration(fromEmail, links)

	for _, link := range links {
		feed := content.GetFeed(link)
		dir := content.MakeMain(feed)

		mobiPath := kindle.GenerateMobi(dir)
		kindle.Send(server, port, fromEmail, password, toEmail, mobiPath)
		kindle.Cleanup(dir)
	}
}
