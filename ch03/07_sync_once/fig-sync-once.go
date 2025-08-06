// fig_sync_once.go
package main

import (
	"fmt"
	"sync"
)

// mainはsync.Onceを使用して関数の実行を一度だけに制限する例を示します。
func main() {
	var count int

	increment := func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)
	for range 100 {
		go func() {
			defer increments.Done()
			once.Do(increment) // 1回しか実行されない
		}()
	}

	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}
