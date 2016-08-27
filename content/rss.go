package content

import (
	"rss-to-kindle/utils"
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
	if len(item.Content) > 0 {
		return item.Content
	}
	if len(item.Extensions["content"]["encoded"]) > 0 && len(item.Extensions["content"]["encoded"][0].Value) > 0 {
		return item.Extensions["content"]["encoded"][0].Value
	}
	return item.Description
}

//GetFeed ...
func GetFeed(path string) Feed {
	feed, err := gofeed.NewParser().ParseURL(path)
	utils.ExitIfErr(err)

	result := Feed{
		Title:       feed.Title,
		Language:    feed.Language,
		Description: feed.Description,
		BuildDate:   *feed.UpdatedParsed,
	}
	result.Sections = append(result.Sections, Section{"Main", nil})

	for key, item := range feed.Items {
		article := Article{
			ID:          key,
			Title:       item.Title,
			Description: item.Description,
			Content:     getContent(*item),
		}
		if item.Author != nil {
			article.Author = item.Author.Name
		}
		result.Sections[0].Articles = append(result.Sections[0].Articles, article)
	}

	return result
}
