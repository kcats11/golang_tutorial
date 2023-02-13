package linuxcmd

import (
	"log"
	"os"
)

func getSourceFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln("getSourceFile error", err)
	}
	defer f.Close()
}

func createDestinationFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalln("getSourceFile error", err)
	}
	defer f.Close()
}

//

func Cp(src string, dest string) {
	getSourceFile(src)
	createDestinationFile(dest)
}
