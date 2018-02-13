package main

import (
    "gopkg.in/fsnotify.v1"
    "log"
)

// ファイルの変更監視
func main(){
    counter := 0
    // 監視用のインスタンス作成
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        panic(err)
    }
    defer watcher.Close()

    done := make(chan bool)
    go func() {
        for {
            select {
            // 変更イベントが入るチャネル(watcher.Events)からイベントを何度も取り出す
            case event := <-watcher.Events:
                log.Println("event:", event)
                if event.Op & fsnotify.Create == fsnotify.Create {
                    log.Println("created file:", event.Name)
                    counter++
                } else if event.Op & fsnotify.Write == fsnotify.Write {
                    log.Println("modified file:", event.Name)
                    counter++
                } else if event.Op & fsnotify.Remove == fsnotify.Remove {
                    log.Println("removed file:", event.Name)
                    counter++
                } else if event.Op & fsnotify.Rename == fsnotify.Rename {
                    log.Println("renamed file:", event.Name)
                    counter++
                } else if event.Op & fsnotify.Chmod== fsnotify.Chmod {
                    log.Println("chmod file:", event.Name)
                    counter++
                }
            case err := <-watcher.Errors:
                log.Println("error:", err)
            }
            if counter > 3 {
                done<-true
            }
        }
    }()

    // 監視対象フォルダを追加する
    err = watcher.Add(".")
    if err != nil {
        log.Fatal(err)
    }
    <-done
}
