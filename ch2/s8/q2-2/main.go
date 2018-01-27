package main

import (
	"os"
	"encoding/csv"
	"log"
)

//  encoding/csvをつかってCSVデータを標準出力とファイルに出力する。
func main() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "Ken"},
		{"Robert", "Griesemer", "gri"},
	}

	// 標準出力のwriter
	w1 := csv.NewWriter(os.Stdout)

	// ファイル出力のwriter
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	w2 := csv.NewWriter(file)

	for _, record := range records {
		if err := w1.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)

		}
		if err := w2.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)

		}
	}
	w1.Flush()
	w2.Flush()
	if err := w1.Error(); err != nil {
		log.Fatal(err)
	}
}

