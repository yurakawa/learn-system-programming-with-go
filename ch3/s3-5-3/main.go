package main

import(
    "encoding/binary"
    "fmt"
    "io"
    "os"
)

func dumpChunk(chunk io.Reader) {
    var length int32
    binary.Read(chunk, binary.BigEndian, &length)
    buffer := make([]byte, 4)
    chunk.Read(buffer)
    fmt.Printf("chunk '%v (%d byte)\n", string(buffer), length)
}

func readChunks(file *os.File) []io.Reader {
    // チャンクを格納する配列
    var chunks []io.Reader

    // 最初の8バイト(シグネチャ)を飛ばす
    file.Seek(8, 0)
    var offset int64 = 8

    for {
        var length int32
        err := binary.Read(file, binary.BigEndian, &length)
        if err == io.EOF {
            break
        }
        chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))

        // 次のチャンクの先頭に移動
        // 現在位置は長さを読み終わった箇所なので
        // チャンク名(4バイト) + データ長 + CRC(4バイト)先に移動
        offset, _ = file.Seek(int64(length+8), 1)
    }
    return chunks
}

func main() {
    // https://upload.wikimedia.org/wikipedia/en/2/24/Lenna.pngより取得
    file, err := os.Open("lenna.png")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    chunks := readChunks(file)
    for _, chunk := range chunks {
        dumpChunk(chunk)
    }

}

