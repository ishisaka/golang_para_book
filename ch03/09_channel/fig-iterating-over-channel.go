package main

import (
	"fmt"
)

// rangeでチャンネルからデータを順次取り出す例
func main() {
	intStream := make(chan int)
	go func() {
		defer close(intStream) // ゴールーチンを抜ける前にチャンネルを閉じる
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream { // rangeでチャンネルからデータを順次取り出す
		fmt.Printf("%v ", integer)
	}
}
