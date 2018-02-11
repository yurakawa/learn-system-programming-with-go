package main

import(
    "fmt"
    "os"
    "syscall"
)

// OS固有のファイル属性を取得する
func main() {
    if len(os.Args) == 1 {
        fmt.Printf("%s [exec file name]", os.Args[0])
        os.Exit(1)
    }
    info, err := os.Stat(os.Args[1])
    if err == os.ErrNotExist {
        fmt.Printf("file not found: %s\n", os.Args[1])
    } else if err != nil {
        panic(err)
    }

    internalStat := info.Sys().(*syscall.Stat_t)
    // macOS固有のファイル属性(Linuxも大差なし)
    fmt.Printf("  デバイス番号 %v\n", internalStat.Dev)
    fmt.Printf("  inode番号 %v\n", internalStat.Ino)
    fmt.Printf("  ブロックサイズ %v\n", internalStat.Blksize)
    fmt.Printf("  ブロック数 %v\n", internalStat.Blocks)
    fmt.Printf("  リンクされている数 %v\n", internalStat.Nlink)
    fmt.Printf("  ファイル作成日時 %v\n", internalStat.Birthtimespec)
    fmt.Printf("  最終アクセス日時 %v\n", internalStat.Atimespec)
    fmt.Printf("  属性変更日時 %v\n", internalStat.Ctimespec)
}
