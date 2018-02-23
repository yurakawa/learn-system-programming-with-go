package main

import (
    "fmt"
    "github.com/AsynkronIT/goconsole"
    "github.com/AsynkronIT/protoactor-go/actor"
)
// アクターモデル

// メッセージの構造体
type hello struct { Who string}
// アクターの構造体
type helloActor struct {}

// アクターのメールボックス受診時に呼ばれるメソッド
func(state *helloActor) Receive(context actor.Context) {
    switch msg := context.Message().(type) {
    case *hello:
        fmt.Printf("Hello %v\n", msg.Who)
    }
}

func main() {
    props := actor.FromInstance(&helloActor{})
    pid := actor.Spawn(props)
    pid.Tell(&hello{Who: "Roger"})
    console.ReadLine()
}
