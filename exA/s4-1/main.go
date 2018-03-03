package main

import (
    "fmt"
    "github.com/tmc/keyring"
    "golang.org/x/crypto/ssh/terminal"
    "syscall"
)

// キーチェーン
func main(){
    secretValue, err := keyring.Get("progo-keyring-test", "password")
    if err == keyring.ErrNotFound {
        // 未登録だった
        fmt.Printf("Secret Value is not found. Please Type:")
        pw, err := terminal.ReadPassword(int(syscall.Stdin))
        if err != nil {
            panic(err)
        }
        // 登録
        // ※ macではapplication passwordとして登録された
        err = keyring.Set("progo-keyring-test", "password", string(pw))
    } else if err != nil {
        // 未知のエラー
        panic(err)
    } else {
        // 登録済みの値を表示
        fmt.Printf("Secret Value: %s\n", secretValue)
    }
}
