package content

import (
	"fmt"
	"regexp"
	"strings"
)

var fileTypeTmpl = `<item href="%s" media-type="%s" id="%s"/>`

func parseArticle(article Article) []string {
	regexTag, _ := regexp.Compile(`<img\s[^>]*?src\s*=\s*['\"]([^'\"]*?)['\"][^>]*?>`)
	regexQuery, _ := regexp.Compile(`\?[\s\S]*?$`)

	result := regexTag.FindAllStringSubmatch(article.Content, -1)

	var images []string
	if len(result) >= 1 {
		for _, item := range result {
			links := item[1]
			links = regexQuery.ReplaceAllString(links, "")
			images = append(images, links)
		}
	}
	return images
}

//GenerateManifest ...
func GenerateManifest(feed Feed) (string, []string) {

	var arr []string
	var images []string

	for _, section := range feed.Sections {
		for _, article := range section.Articles {
			articleImages := parseArticle(article)
			images = append(images, articleImages...)

			idStr := fmt.Sprintf("%d", article.ID)
			str := fmt.Sprintf(fileTypeTmpl, idStr+".html", "application/xhtml+xml", idStr)
			arr = append(arr, str)
		}
	}
	arr = append(arr,
		fmt.Sprintf(fileTypeTmpl, "contents.html", "application/xhtml+xml", "contents"),
		fmt.Sprintf(fileTypeTmpl, "nav-contents.ncx", "application/x-dtbncx+xml", "nav-contents"),
	)
	return strings.Join(arr, "\n"), images
}
