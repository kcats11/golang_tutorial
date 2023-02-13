package linuxcmd

import (
	"fmt"
	"log"
	"os"
)

func Ls(dir string) {
	var files []os.DirEntry
	var err error

	if len(dir) > 0 {
		files, err = os.ReadDir(dir)
	} else {
		files, err = os.ReadDir(".")
	}

	if err != nil {
		log.Fatalln("ls error :", err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
