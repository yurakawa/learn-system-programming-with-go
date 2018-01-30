package main

import (
	"net/http"
	"io"
	"archive/zip"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment;filename=a.zip")

	//file, err := os.Create("ascii_sample.zip")
	//if err != nil {
	//	panic(err)
	//}
	// writer作成
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	a, err := zipWriter.Create("a.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(a, strings.NewReader(("Q3.4: zipファイルをWebサーバからダウンロード")))
}

// zipファイルをWebサーバからダウンロード
func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}


