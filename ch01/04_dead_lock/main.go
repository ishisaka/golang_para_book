/**

1.2.4 デッドロックの例

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	type value struct {
		mu    sync.Mutex
		value int
	}

	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()         // ❶ クリティカルセクションに入る
		defer v1.mu.Unlock() // ❷ deferでロックを解放
		// ❸ 処理の負荷をシミュレートするために
		// スリープするがここでデッドロックが発生する。
		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
