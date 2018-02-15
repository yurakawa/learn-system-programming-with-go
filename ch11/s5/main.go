package main

import (
    "fmt"
    "os"
    "os/exec"
)

// exec.Cmdによるプロセスの起動
// 引数に外部プログラムを指定すると、実行にかかった時間を表示する
func main() {
    if len(os.Args) == 1 {
        return
    }
    cmd := exec.Command(os.Args[1], os.Args[2:]...)
    err := cmd.Run()
    if err != nil {
        panic(err)
    }
    state := cmd.ProcessState
    fmt.Printf("%s\n", state.String())
    fmt.Printf("  Pid: %d\n", state.Pid())
    fmt.Printf("  Exited: %v\n", state.Exited())
    fmt.Printf("  Success: %v\n", state.Success())
    fmt.Printf("  System: %v\n", state.SystemTime())
    fmt.Printf("  User:  %d\n", state.UserTime())
}
