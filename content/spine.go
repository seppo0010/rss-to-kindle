package content

import (
	"fmt"
	"strings"
)

//GenerateSpine ...
func GenerateSpine(feed Feed) string {
	var arr []string
	arr = append(arr, fmt.Sprintf(SpineTmpl, "contents"))

	for _, section := range feed.Sections {
		for _, article := range section.Articles {
			arr = append(arr, fmt.Sprintf(SpineTmpl, fmt.Sprintf("%d", article.ID)))
		}
	}

	return strings.Join(arr, "\n")
}
