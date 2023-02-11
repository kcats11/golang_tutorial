package main

import (
	"log"
	"os"
)

// ファイル1の内容を取得
func getSourceFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln("getSourceFile error", err)
	}
	defer f.Close()
}

// 取得したファイルの中身をもファイル2を作成
func createDestinationFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalln("getSourceFile error", err)
	}
	defer f.Close()
}

//

func main() {
	src, dest := os.Args[1], os.Args[2]
	getSourceFile(src)
	createDestinationFile(dest)
}
