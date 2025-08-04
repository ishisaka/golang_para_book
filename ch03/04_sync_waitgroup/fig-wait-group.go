// fig-wait-group.go
// Goルーチンの同期をWaitGroupで行う例
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1) // WaitGroupに1つのゴルーチンを追加
	go func() {
		defer wg.Done() // ゴルーチンの終了時にWaitGroupから1つ減らす
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()

	wg.Add(1) // WaitGroupに1つのゴルーチンを追加
	go func() {
		defer wg.Done() // ゴルーチンの終了時にWaitGroupから1つ減らす
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait() // 全てのゴルーチンが終了するのを待つ
	fmt.Println("All goroutines complete.")
}
