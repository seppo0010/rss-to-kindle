package content

import (
	"fmt"
)

var articleTmpl = `<html>
	<head>
		<meta content="text/html; charset=utf-8" http-equiv="Content-Type"/>
		<title>%s</title>
	</head>
	<body>
		%s
	</body>
</html>`

//GenerateArticle ...
func GenerateArticle(article Article) string {
	return fmt.Sprintf(articleTmpl, article.Title, article.Content)
}
