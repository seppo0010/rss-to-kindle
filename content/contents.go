package content

import (
	"fmt"
	"strings"
)

//GenerateContents ...
func GenerateContents(feed Feed) string {
	var sectionArr []string
	for _, section := range feed.Sections {
		sectionArr = append(sectionArr, generateSection(section))
	}
	return fmt.Sprintf(ContentPageTmpl, strings.Join(sectionArr, "\n"))
}

func generateSection(section Section) string {
	var articleArr []string
	for _, article := range section.Articles {
		articleArr = append(articleArr, generateArticle(article))
	}
	return fmt.Sprintf(ContentSectionTmpl, section.Title, strings.Join(articleArr, "\n"))
}

func generateArticle(article Article) string {
	return fmt.Sprintf(ContentArticleTmpl, fmt.Sprintf("%d.html", article.ID), article.Title)
}
