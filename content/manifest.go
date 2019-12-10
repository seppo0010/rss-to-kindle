package content

import (
	"fmt"
	"regexp"
	"github.com/seppo0010/rss-to-kindle/utils"
	"strings"
)

func parseArticle(article Article) []string {
	var images []string

	result := regexp.MustCompile(ImageTagRegex).FindAllStringSubmatch(article.Content, -1)
	for _, item := range result {
		images = append(images, utils.StripQuery(item[1]))
	}

	return images
}

//GenerateManifest ...
func GenerateManifest(feed Feed) (string, []string) {
	var arr []string
	var images []string

	for _, section := range feed.Sections {
		for _, article := range section.Articles {
			images = append(images, parseArticle(article)...)
			idStr := fmt.Sprintf("%d", article.ID)
			arr = append(arr, fmt.Sprintf(ManifestItemTmpl, idStr+".html", "application/xhtml+xml", idStr))
		}
	}

	arr = append(arr,
		fmt.Sprintf(ManifestItemTmpl, "contents.html", "application/xhtml+xml", "contents"),
		fmt.Sprintf(ManifestItemTmpl, "nav-contents.ncx", "application/x-dtbncx+xml", "nav-contents"),
	)

	return strings.Join(arr, "\n"), images
}
