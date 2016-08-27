package utils

import "log"

//ExitIfErr ...
func ExitIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
