package kindle

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"github.com/seppo0010/rss-to-kindle/utils"
)

//GenerateMobi ...
func GenerateMobi(dir string) string {
	cmd := exec.Command("kindlegen", filepath.Join(dir, "main.opf"), "-c2", "-gif")
	cmd.Run()

	mobiPath := filepath.Join(dir, "main.mobi")
	_, err := ioutil.ReadFile(mobiPath)
	utils.ExitIfErr(err)

	return mobiPath
}
