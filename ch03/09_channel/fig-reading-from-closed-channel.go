package main

import (
	"fmt"
)

// 閉じられたChannelから値を読み込んだ場合の例
func main() {
	intStream := make(chan int)
	close(intStream)
	integer, ok := <-intStream          // 閉じられたChannelから値を読み込む
	fmt.Printf("(%v): %v", ok, integer) // (false): 0
}
