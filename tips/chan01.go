/********************
 * chan01.go
 * 2016/12/12(月)
 * walkingmask
 * channel sample 1
 ********************/

package main

import "fmt"

func main() {
  n := make(chan int, 0)

// ここに入れるとエラー
// nはbufferなしだから送信側は受信操作が行われるまでsleepする
// main threadがsleepするのはout
// 送信開始 -> 受信開始 -> 送信 -> 受信 -> 送信側起こす (同期)
// n <- 2
  go func() {
    fmt.Println(<-n)
    n <- 1
  }()
  n <- 2

  fmt.Println(<-n)
}

/*
その他チャネルについてのメモ
 * eventのやり取りだけしたchanにはstruct{}を使う
 * これ以上メッセージが通信されないことを明示 -> close
 * closeしなくても到達不可能になったchanはGCが回収
 * v, ok := <- ch でcloseされたchanからfalseが来る
 : for v := range ch でも同じ
*/
