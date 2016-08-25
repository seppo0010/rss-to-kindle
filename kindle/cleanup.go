package kindle

import "os"

//Cleanup ...
func Cleanup(dir string) {
	os.RemoveAll(dir)
	//fmt.Println(dir)
}
