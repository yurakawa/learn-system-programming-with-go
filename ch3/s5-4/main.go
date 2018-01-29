package main

import(
    "encoding/binary"
    "io"
    "os"
    "bytes"
    "hash/crc32"
)

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

func textChunk(text string) io.Reader {
    byteData := []byte(text)
    var buffer bytes.Buffer
    binary.Write(&buffer, binary.BigEndian, int32(len(byteData)))
    buffer.WriteString("tEXt")
    buffer.Write(byteData)
    // CRCを計算して追加
    crc := crc32.NewIEEE()
    io.WriteString(crc, "tEXt")
    binary.Write(&buffer, binary.BigEndian, crc.Sum32())
    return &buffer
}

func main() {
    // https://upload.wikimedia.org/wikipedia/en/2/24/Lenna.pngより取得
    file, err := os.Open("lenna.png")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    newFile, err := os.Create("Lenna2.png")
    if err != nil {
        panic(err)
    }
    defer newFile.Close()
    chunks := readChunks(file)
    // シグニチャ書き込み
    io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
    // 先頭に必要なIHDRチャンクを書き込み
    io.Copy(newFile, chunks[0])
    // テキストチャンクを追加
    io.Copy(newFile, textChunk("ASCII PROGRAMMING++"))
    // 残りのチャンクを追加
    for _, chunk := range chunks {
        io.Copy(newFile, chunk)
    }

}

