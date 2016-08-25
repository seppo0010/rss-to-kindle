package kindle

import (
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
)

//GenerateMobi ...
func GenerateMobi(dir string) string {
	cmd := exec.Command("kindlegen", filepath.Join(dir, "main.opf"), "-c2", "-gif")

	cmd.Run()

	mobiPath := filepath.Join(dir, "main.mobi")
	_, err := ioutil.ReadFile(mobiPath)
	if err != nil {
		log.Fatal(err)
	}

	return mobiPath
}
