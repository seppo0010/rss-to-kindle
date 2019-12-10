package content

import "fmt"

//GenerateOpf ...
func GenerateOpf(feed Feed, manifest string, spine string) string {
	creator := feed.Title
	publisher := feed.Title
	if feed.Language == "" { feed.Language = "en" }
	return fmt.Sprintf(OpfTmpl,
		feed.ID,
		feed.Title,
		feed.Language,
		creator,
		publisher,
		feed.BuildDate.Format("2006-01-02"),
		feed.Description,
		manifest,
		spine,
	)
}
