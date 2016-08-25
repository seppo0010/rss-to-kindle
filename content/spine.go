package content

import (
	"fmt"
	"strings"
)

var spineTmpl = `<itemref idref="%s"/>`

//GenerateSpine ...
func GenerateSpine(feed Feed) string {
	var arr []string
	arr = append(arr, fmt.Sprintf(spineTmpl, "contents"))
	for _, section := range feed.Sections {
		for _, article := range section.Articles {
			idStr := fmt.Sprintf("%d", article.ID)
			str := fmt.Sprintf(spineTmpl, idStr)
			arr = append(arr, str)
		}
	}
	return strings.Join(arr, "\n")
}
