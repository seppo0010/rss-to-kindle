package content

import (
	"fmt"
	"strings"
)

func generateNavHeader(feed Feed) string {
	var navHeaderTmpl = `<head>
		<meta content="%s" name="dtb:uid"/>
		<meta content="2" name="dtb:depth"/>
		<meta content="0" name="dtb:totalPageCount"/>
		<meta content="0" name="dtb:maxPageNumber"/>
	</head>
	<docTitle>
		<text>%s/text>
	</docTitle>
	<docAuthor>
		<text>$s</text>
	</docAuthor>`
	return fmt.Sprintf(navHeaderTmpl,
		feed.ID,
		feed.Title,
		feed.Title,
	)
}

func generateNavSections(feed Feed) string {
	var currentPlayOrder = 1
	var navSectionTmpl = `<navPoint playOrder="%d" class="section" id="%s">
		<navLabel>
			<text>%s</text>
		</navLabel>
		<content src="%s"/>
		%s
	</navPoint>`

	var sectionsArr []string
	for _, section := range feed.Sections {
		articleStr, newPlayOrder := generateNavArticles(section.Articles, currentPlayOrder)

		content := fmt.Sprintf(navSectionTmpl,
			currentPlayOrder,
			section.Title,
			section.Title,
			fmt.Sprintf("%d.html", section.Articles[0].ID),
			articleStr,
		)

		currentPlayOrder = newPlayOrder + 1
		sectionsArr = append(sectionsArr, content)
	}
	return strings.Join(sectionsArr, "\n")
}

func generateNavArticles(articles []Article, currentPlayOrder int) (string, int) {
	var navArticleTmpl = `<navPoint playOrder="%d" class="article" id="%s">
		<navLabel>
			<text>%s</text>
		</navLabel>
		<content src="%s"/>
		<mbp:meta name="description">%s</mbp:meta>
		<mbp:meta name="author">%s</mbp:meta>
	</navPoint>
	`
	var articleArr []string
	for _, article := range articles {
		currentPlayOrder++
		content := fmt.Sprintf(navArticleTmpl,
			currentPlayOrder,
			article.ID,
			article.Title,
			fmt.Sprintf("%d.html", article.ID),
			article.Description,
			article.Author,
		)
		articleArr = append(articleArr, content)
	}
	return strings.Join(articleArr, "\n"), currentPlayOrder
}

func generateNavMap(feed Feed) string {
	var navMapTmpl = `<navMap>
		<navPoint playOrder="0" class="periodical" id="periodical">
			<navLabel>
				<text>Table of Contents</text>
			</navLabel>
			<content src="contents.html"/>
			%s
		</navPoint>
	</navMap>`

	return fmt.Sprintf(navMapTmpl, generateNavSections(feed))
}

//GenerateNavMain ...
func GenerateNavMain(feed Feed) string {
	var navMainTmpl = `<?xml version='1.0' encoding='utf-8'?>
		<!DOCTYPE ncx PUBLIC "-//NISO//DTD ncx 2005-1//EN" "http://www.daisy.org/z3986/2005/ncx-2005-1.dtd">
		<ncx xmlns:mbp="http://mobipocket.com/ns/mbp" xmlns="http://www.daisy.org/z3986/2005/ncx/" version="2005-1" xml:lang="en-GB">
			%s
			%s
		</ncx>`

	return fmt.Sprintf(navMainTmpl,
		generateNavHeader(feed),
		generateNavMap(feed),
	)
}
