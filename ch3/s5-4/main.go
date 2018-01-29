package main

import (
    "io"
    "encoding/binary"
    "fmt"
    "bytes"
    "os"
    "hash/crc32"
)

func dumpChunk(chunk io.Reader) {
    var l int32
    binary.Read(chunk, binary.BigEndian, &l)
    // png をbyteに変換した時の入れ物
    buf := make([]byte, 4)
    chunk.Read(buf)
    fmt.Printf("chunk '%v' (%d byte)\n", string(buf), l)
    if bytes.Equal(buf, []byte("tExt")) {
        rawText := make([]byte, l)
        chunk.Read(rawText)
        fmt.Printf("%s\n", string(rawText))
    }
}

func readChunks(file *os.File) []io.Reader {
    // チャンクを格納する配列
    var chunks []io.Reader

    // pngの最初の8byte(シグネチャ)を飛ばして9byte目から読み込む
    file.Seek(8, 0)
    var ofs int64 = 8

    for {
        var length int32
        err := binary.Read(file, binary.BigEndian, &length)
        if err == io.EOF {
            break
        }

        chunks = append(chunks, io.NewSectionReader(file, ofs, int64(length)+12))

        // 次のチャンクの先頭に移動
        // 現在位置は長さを読み終わった箇所なので
        // チャンク名(4バイト) + データ長 + CRC(4バイト)先に移動
        ofs, _ = file.Seek(int64(length+8), 1)
    }

    return chunks
}

func textChunk(txt string) io.Reader {
    byteData := []byte(txt)
    var buffer bytes.Buffer
    binary.Write(&buffer, binary.BigEndian, int32(len(byteData))) // 書き込むためのbufferを確保。長さは入力文字長。
    buffer.WriteString("tExt")                                    //bufferに書き込む
    buffer.Write(byteData)                                        // 入力対象の文字をbufferに書き込む
    // CRCを計算して追加
    crc := crc32.NewIEEE()
    io.WriteString(crc, "tExt")
    binary.Write(&buffer, binary.BigEndian, crc.Sum32()) // crc分を新しくbufferに書き込む
    return &buffer                                       // 書き込み終わったbufferを返す
}

func WriteNewTextChunk() {
    // pngをチャンクに分けて読み込み
    file, err := os.Open("Lenna.png")
    if err != nil {
        fmt.Printf("error! err: %v", err)
        os.Exit(-1)
    }
    defer file.Close()

    newFile, _ := os.Create("output.png")
    defer newFile.Close()

    chunks := readChunks(file)
    // pngのシグニチャー書き込み
    io.WriteString(newFile, "\x89PNG\r\n\x1a\n")

    // 先頭にIHDR chunkを書き込む
    io.Copy(newFile, chunks[0])

    // TextChunk を追加する
    io.Copy(newFile, textChunk("Test Text Chunk"))


    // 残りのチャンクを新しく追加する
    for _, c := range chunks[1:] {
        io.Copy(newFile, c)
    }
}

func main() {
    // chunkを書き込んだfileを生成する
    WriteNewTextChunk()

    file, err := os.Open("output.png")
    if err != nil {
        fmt.Printf("error! err: %v", err)
        os.Exit(-1)
    }

    chunks := readChunks(file)
    for _, c := range chunks {
        dumpChunk(c)
    }
}