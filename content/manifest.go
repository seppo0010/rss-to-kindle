package content

import (
	"fmt"
	"strings"
)

var fileTypeTmpl = `<item href="%s" media-type="%s" id="%s"/>`

//GenerateManifest ...
func GenerateManifest(feed Feed) string {
	var arr []string
	for _, section := range feed.Sections {
		for _, article := range section.Articles {
			idStr := fmt.Sprintf("%d", article.ID)
			fileName := idStr + ".html"
			fileType := "application/xhtml+xml"
			str := fmt.Sprintf(fileTypeTmpl, fileName, fileType, idStr)
			arr = append(arr, str)
		}
	}
	return strings.Join(arr, "\n")
}
