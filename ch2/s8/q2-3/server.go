package main


import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"log"
)

// gzipされたJSONをクライアントに返しながら標準出力にも出力する
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contet-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json化する元のデータ
	source := map[string]string{
		"Hello": "world",
	}
	g := gzip.NewWriter(w)
	out := io.MultiWriter(g, os.Stdout)
	encoder := json.NewEncoder(out)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(source); err != nil {
		log.Print(err)
	}
	if err := g.Close(); err != nil {
		log.Print(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

