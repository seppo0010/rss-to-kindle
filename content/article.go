package content

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

//GenerateArticle ...
func GenerateArticle(article Article) string {
	var articleTmpl = `<html>
		<head>
			<meta content="text/html; charset=utf-8" http-equiv="Content-Type"/>
			<title>%s</title>
		</head>
		<body>
			<div class="header">
				<h1>%s</h1>
			</div>
			%s
		</body>
	</html>`

	var imageTagTmpl = `<p><img src="%s.jpg" middle="true"></p>`

	r, _ := regexp.Compile(`<img[\s\S]+src="(?P<src>[\s\S]*?)"[\s\S]*\/>`)
	result := r.FindAllStringSubmatch(article.Content, -1)

	if len(result) >= 1 {
		for _, item := range result {
			_, filename := filepath.Split(item[1])
			filename = strings.TrimSuffix(filename, filepath.Ext(filename))

			imageTag := fmt.Sprintf(imageTagTmpl, filename)
			article.Content = strings.Replace(article.Content, item[0], imageTag, -1)
			splitedContent := strings.Split(article.Content, "\n")

			for key, paragraph := range splitedContent {
				paragraph = strings.Replace(paragraph, "<br>", "", -1)
				paragraph = strings.TrimSpace(paragraph)

				if len(paragraph) > 0 {
					splitedContent[key] = "<p>" + paragraph + "</p><br>"
				}
			}

			for key, paragraph := range splitedContent {
				if paragraph == "" {
					if key+1 >= len(splitedContent) {
						splitedContent = splitedContent[:key]
					} else {
						splitedContent = append(splitedContent[:key], splitedContent[key+1:]...)
					}
				}
			}

			article.Content = strings.Join(splitedContent, "\n")
		}
	}

	return fmt.Sprintf(articleTmpl, article.Title, article.Title, article.Content)
}
