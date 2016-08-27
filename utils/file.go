package utils

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

//StripQuery ...
func StripQuery(path string) string {
	return regexp.MustCompile(`\?[\s\S]*?$`).ReplaceAllString(path, "")
}

//GetFilename ...
func GetFilename(path string, includeExt bool) string {
	_, filename := filepath.Split(StripQuery(path))
	if includeExt == true {
		return filename
	}
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

//WriteFile ...
func WriteFile(dir string, filename string, content []byte) {
	path := filepath.Join(dir, filename)
	err := ioutil.WriteFile(path, content, 0666)
	ExitIfErr(err)
}

//NormalizeImageFilename ...
func NormalizeImageFilename(filename string) string {
	filename = strings.Replace(filename, ".PNG", ".png", -1)
	filename = strings.Replace(filename, ".JPEG", ".jpg", -1)
	filename = strings.Replace(filename, ".JPG", ".jpg", -1)
	filename = strings.Replace(filename, ".jpeg", ".jpg", -1)
	return filename
}
