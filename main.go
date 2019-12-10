package main

import (
	"fmt"
	"log"
	"os"
	"github.com/seppo0010/rss-to-kindle/content"
	"github.com/seppo0010/rss-to-kindle/kindle"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type configuration struct {
	server    string
	port      string
	fromEmail string
	password  string
	toEmail   string
	links     []string
}

func printHeader(message string) {
	c := color.New(color.FgCyan)

	c.Println("---")
	c.Println(message)
	c.Println("---")
}

func printStatus(message string) {
	c := color.New(color.FgGreen)
	c.Println(message)
}

func printConfiguration(conf configuration) {
	fmt.Println("SERVER:\t" + conf.server)
	fmt.Println("PORT:\t" + conf.port)
	fmt.Println("FROM:\t" + conf.fromEmail)
	fmt.Println("PWD:\t" + "[redacted]")
	fmt.Println("TO:\t" + conf.toEmail)
	if len(conf.links) > 0 {
		fmt.Printf("LINKS:")
		for _, link := range conf.links {
			fmt.Println("\t* " + link)
		}
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := configuration{
		os.Getenv("SERVER"),
		os.Getenv("PORT"),
		os.Getenv("FROM_EMAIL"),
		os.Getenv("PASSWORD"),
		os.Getenv("TO_EMAIL"),
		strings.Split(os.Getenv("LINKS"), ","),
	}

	printHeader("Configuration")
	printConfiguration(conf)

	for _, link := range conf.links {
		feed := content.GetFeed(link)
		printHeader("Feed: " + feed.Title)

		printStatus("Creating files...")
		dir := content.MakeMain(feed)

		printStatus("Generating .mobi file...")
		mobiPath := kindle.GenerateMobi(dir)
		printStatus("path: " + mobiPath)

		os.Rename(mobiPath, "./target.mobi")
		printStatus("Cleaning up...")
		kindle.Cleanup(dir)

		printStatus("Done.")
	}
}
