// fig-bulk-add.go
// Goルーチンの同期をWaitGroupで行う例
package main

import (
	"fmt"
	"sync"
)

func main() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}

	const numGreeters = 5
	var wg sync.WaitGroup
	// WaitGroupに複数のゴルーチンを一度に追加
	// Addメソッドを使って、ゴルーチンの数を指定する
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}
