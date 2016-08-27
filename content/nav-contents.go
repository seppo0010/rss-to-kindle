package content

import (
	"fmt"
	"strings"
)

func generateNavHeader(feed Feed) string {
	return fmt.Sprintf(NavHeaderTmpl, feed.ID, feed.Title, feed.Title)
}

func generateNavArticles(articles []Article, currentPlayOrder int) (string, int) {
	var articleArr []string

	for _, article := range articles {
		currentPlayOrder++
		content := fmt.Sprintf(NavArticleTmpl,
			currentPlayOrder,
			article.ID,
			article.Title,
			fmt.Sprintf("%d.html", article.ID),
			article.Description,
			article.Author,
		)
		articleArr = append(articleArr, content)
	}

	return strings.Join(articleArr, "\n"), currentPlayOrder
}

func generateNavSections(feed Feed) string {
	var sectionsArr []string

	var currentPlayOrder = 1
	for _, section := range feed.Sections {
		articleStr, newPlayOrder := generateNavArticles(section.Articles, currentPlayOrder)
		content := fmt.Sprintf(NavSectionTmpl,
			currentPlayOrder,
			section.Title,
			section.Title,
			fmt.Sprintf("%d.html", section.Articles[0].ID),
			articleStr,
		)
		currentPlayOrder = newPlayOrder + 1
		sectionsArr = append(sectionsArr, content)
	}

	return strings.Join(sectionsArr, "\n")
}

func generateNavMap(feed Feed) string {
	return fmt.Sprintf(NavMapTmpl, generateNavSections(feed))
}

//GenerateNavMain ...
func GenerateNavMain(feed Feed) string {
	return fmt.Sprintf(NavMainTmpl, generateNavHeader(feed), generateNavMap(feed))
}
