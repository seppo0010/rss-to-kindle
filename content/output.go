package content

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/nfnt/resize"
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

func resizeImageFile(path string) {
	r, _ := regexp.Compile(`.jpg$|.png$`)
	if r.MatchString(path) == false {
		return
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	var img image.Image
	ext := filepath.Ext(path)
	if ext == ".jpg" {
		img, err = jpeg.Decode(file)
	} else if ext == ".png" {
		img, err = png.Decode(file)
	}

	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	os.Remove(path)

	m := resize.Resize(640, 0, img, resize.Bicubic)

	newPath := strings.Replace(path, ".png", ".jpg", -1)

	out, err := os.Create(newPath)

	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	err = jpeg.Encode(out, m, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func normalizeImageFilename(filename string) string {
	filename = strings.Replace(filename, ".PNG", ".png", -1)
	filename = strings.Replace(filename, ".JPEG", ".jpg", -1)
	filename = strings.Replace(filename, ".JPG", ".jpg", -1)
	filename = strings.Replace(filename, ".jpeg", ".jpg", -1)
	return filename
}

func makeImageFile(dir string, images []string) {
	for _, image := range images {
		res, err := http.Get(image)
		if err != nil {
			log.Fatal(err)
		}
		content, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		_, filename := filepath.Split(image)
		filename = normalizeImageFilename(filename)
		if err = ioutil.WriteFile(filepath.Join(dir, filename), content, 0666); err != nil {
			log.Fatal(err)
		}
		resizeImageFile(filepath.Join(dir, filename))
	}
}

func makeOpfFile(dir string, feed Feed) []string {
	path := filepath.Join(dir, "main.opf")

	manifest, images := GenerateManifest(feed)
	spine := GenerateSpine(feed)

	makeImageFile(dir, images)

	content := []byte(GenerateOpf(feed, manifest, spine))
	if err := ioutil.WriteFile(path, content, 0666); err != nil {
		log.Fatal(err)
	}

	return images
}

func makeNavContentsFile(dir string, feed Feed) {
	path := filepath.Join(dir, "nav-contents.ncx")

	content := []byte(GenerateNavMain(feed))
	if err := ioutil.WriteFile(path, content, 0666); err != nil {
		log.Fatal(err)
	}
}

//MakeMain ...
func MakeMain(feed Feed) string {
	dir := createTempDir()

	makeContentsFile(dir, feed)
	makeOpfFile(dir, feed)
	makeArticlesFile(dir, feed)
	makeNavContentsFile(dir, feed)

	return dir
}
