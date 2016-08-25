package content

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func createTempDir() string {
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		fmt.Println(err)
	}
	return dir
}

func makeArticlesFile(dir string, feed Feed) {
	for _, section := range feed.Sections {
		for _, article := range section.Articles {
			path := filepath.Join(dir, fmt.Sprintf("%d.html", article.ID))
			content := []byte(GenerateArticle(article))
			if err := ioutil.WriteFile(path, content, 0666); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func makeContentsFile(dir string, feed Feed) {
	path := filepath.Join(dir, "contents.html")
	content := []byte(GenerateContents(feed))
	if err := ioutil.WriteFile(path, content, 0666); err != nil {
		log.Fatal(err)
	}
}

func makeOpfFile(dir string, feed Feed) {
	path := filepath.Join(dir, "main.opf")

	manifest := GenerateManifest(feed)
	spine := GenerateSpine(feed)

	content := []byte(GenerateOpf(feed, manifest, spine))
	if err := ioutil.WriteFile(path, content, 0666); err != nil {
		log.Fatal(err)
	}
}

//MakeMain ...
func MakeMain(feed Feed) string {
	dir := createTempDir()
	fmt.Println(dir)

	makeContentsFile(dir, feed)
	makeArticlesFile(dir, feed)
	makeOpfFile(dir, feed)

	return dir
}
