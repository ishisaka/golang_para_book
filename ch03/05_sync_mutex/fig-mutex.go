// fig-mutex.go
// このコードは、複数のゴルーチンが共有変数 count を更新する際に、sync.Mutex を
// 用いてデータ競合を防ぎ、安全に処理を行うサンプルプログラムです。
package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()         // クリティカルセクションの開始
		defer lock.Unlock() // deferでクリティカルセクションの終了を保証
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()         // クリティカルセクションの開始
		defer lock.Unlock() // deferでクリティカルセクションの終了を保証
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	// Increment
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	// Decrement
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")
}
