package content

import (
	"fmt"
	"regexp"
	"rss-to-kindle/utils"
	"strings"
)

//GenerateArticle ...
func GenerateArticle(article Article) string {
	result := regexp.MustCompile(ImageTagRegex).FindAllStringSubmatch(article.Content, -1)

	for _, item := range result {
		imageTag := fmt.Sprintf(ImageTagTmpl, utils.GetFilename(item[1], false))

		article.Content = strings.Replace(article.Content, item[0], imageTag, -1)
		splitedContent := strings.Split(article.Content, "\n")

		for key, paragraph := range splitedContent {
			paragraph = strings.TrimSpace(strings.Replace(paragraph, "<br>", "", -1))
			if len(paragraph) > 0 {
				splitedContent[key] = "<p>" + paragraph + "</p>"
			}
		}

		for key, paragraph := range splitedContent {
			if len(paragraph) < 0 {
				if key+1 >= len(splitedContent) {
					splitedContent = splitedContent[:key]
				} else {
					splitedContent = append(splitedContent[:key], splitedContent[key+1:]...)
				}
			}
		}

		article.Content = strings.Join(splitedContent, "\n")
	}

	return fmt.Sprintf(ArticleTmpl, article.Title, article.Title, article.Content)
}
