package main

import (
	"fmt"
)

// 基本的なChannelの使用方法
func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!" // チャンネルにデータを送る
	}()
	fmt.Println(<-stringStream) // チャンネルからデータを受け取る
}
