package main

import (
	"fmt"
)

// Channelは自身の状態を返すことができる
func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()
	salutation, ok := <-stringStream // 2つ目の戻り津はChannelの状態を返す
	fmt.Printf("(%v): %v", ok, salutation)
}
