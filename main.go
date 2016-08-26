package main

import (
	"fmt"
	"log"
	"os"
	"rss-to-kindle/content"
	"rss-to-kindle/kindle"
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

func printConfiguration(conf configuration) {
	color.New(color.FgCyan).Add(color.Underline).Println("Configuration")
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

	printConfiguration(conf)

	for _, link := range conf.links {
		feed := content.GetFeed(link)
		dir := content.MakeMain(feed)

		mobiPath := kindle.GenerateMobi(dir)

		kindle.Send(
			conf.server,
			conf.port,
			conf.fromEmail,
			conf.password,
			conf.toEmail,
			mobiPath,
		)

		kindle.Cleanup(dir)
	}
}
