package content

import (
	"fmt"
	"regexp"
	"strings"
)

var fileTypeTmpl = `<item href="%s" media-type="%s" id="%s"/>`

func parseArticle(article Article) []string {
	r, _ := regexp.Compile(`<img[\s\S]+src="(?P<src>[\s\S]*?)"[\s\S]*\/>`)
	result := r.FindAllStringSubmatch(article.Content, -1)

	var images []string

	if len(result) >= 1 {
		for _, item := range result {
			links := item[1]
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
	return strings.Join(arr, "\n"), images
}
