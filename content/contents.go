package content

import (
	"fmt"
	"strings"
)

//GenerateContents ...
func GenerateContents(feed Feed) string {
	var mainTmpl = `<html>
		<head>
			<meta content="text/html; charset=utf-8" http-equiv="Content-Type"/>
			<title>Table of Contents</title>
		</head>
		<body>
			<h1>Contents</h1>
			%s
		</body>
	</html>`
	var sectionArr []string
	for _, section := range feed.Sections {
		sectionArr = append(sectionArr, generateSection(section))
	}
	sectionStr := strings.Join(sectionArr, "\n")
	return fmt.Sprintf(mainTmpl, sectionStr)
}

func generateSection(section Section) string {
	var sectionTmpl = `<h4>%s</h4>
	<ul>
		%s
	</ul>`
	var articleArr []string
	for _, article := range section.Articles {
		articleArr = append(articleArr, generateArticle(article))
	}
	articleStr := strings.Join(articleArr, "\n")
	return fmt.Sprintf(sectionTmpl, section.Title, articleStr)
}

func generateArticle(article Article) string {
	var articleTmpl = `<li>
		<a href="%s">%s</a>
	</li>`
	path := fmt.Sprintf("%d.html", article.ID)
	return fmt.Sprintf(articleTmpl, path, article.Title)
}
