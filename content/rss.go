package content

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

//Feed ...
type Feed struct {
	ID          string
	Title       string
	Language    string
	Description string
	BuildDate   time.Time
	Sections    []Section
}

//Section ...
type Section struct {
	Title    string
	Articles []Article
}

//Article ...
type Article struct {
	ID          int
	Title       string
	Description string
	Content     string
	Author      string
}

func getContent(item gofeed.Item) string {
	if item.Content != "" {
		return item.Content
	}
	if len(item.Extensions["content"]["encoded"]) > 0 && len(item.Extensions["content"]["encoded"][0].Value) > 0 {
		return item.Extensions["content"]["encoded"][0].Value
	}
	return item.Description
}

//GetFeed ...
func GetFeed(path string) Feed {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(path)

	if err != nil {
		fmt.Println("Error retriving feed with " + path + ".")
	}

	result := Feed{}

	result.Title = feed.Title
	result.Language = feed.Language
	result.Description = feed.Description
	result.BuildDate = *feed.UpdatedParsed
	result.Sections = append(result.Sections, Section{"Main", nil})

	for key, item := range feed.Items {

		article := Article{}
		article.ID = key
		article.Title = item.Title
		article.Description = item.Description

		article.Content = getContent(*item)

		if item.Author != nil {
			article.Author = item.Author.Name
		}

		result.Sections[0].Articles = append(result.Sections[0].Articles, article)
	}
	return result
}
