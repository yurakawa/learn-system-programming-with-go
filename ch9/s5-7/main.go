package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

var imageSuffix = map[string]bool{
    ".jpg":  true,
    ".jpeg": true,
    ".png":  true,
    ".webp": true,
    ".gif":  true,
    ".tiff": true,
    ".eps":  true,
}

// ディレクトリのトラバース
func main() {
    if len(os.Args) == 1 {
        fmt.Printf(`Find images
 
Usage:
    %s [path to find]
`, os.Args[0])
        return
    }
    root := os.Args[1]

    err := filepath.Walk(root,
        func(path string, info os.FileInfo, err error) error {
            // ディレクトリのときは何もしない
            if info.IsDir() {
                // Sphinx用関連なのでSkip
                if info.Name() == "_build" {
                    return filepath.SkipDir
                }
                return nil
            }
            ext := strings.ToLower(filepath.Ext(info.Name()))
            if imageSuffix[ext] {
                // 表記が煩雑にならないようにする.
                // ファイルの内容を読み込むようなときはos.Open()にそのまま渡せる
                rel, err := filepath.Rel(root, path)
                if err != nil {
                    return nil
                }
                fmt.Printf("%s\n", rel)
            }
            return nil
        })
    if err != nil {
        fmt.Println(1, err)
    }
}

