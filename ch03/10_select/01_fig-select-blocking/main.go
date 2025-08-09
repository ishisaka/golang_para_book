// main.go
package main

import (
	"fmt"
	"time"
)

// main 関数は、ゴルーチンで非同期処理を実行し、チャネルのクローズを監視してブロックの解除を行います。
func main() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c) // 5秒待ってチャネルを閉じる
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c: // チャネルを読み込む。
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}
