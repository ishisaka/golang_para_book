// fig_sync_once_diff_funcs.go
package main

import (
	"fmt"
	"sync"
)

// mainはsync.Onceを使用して関数の実行を1回に制限する例を示します。
func main() {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decrement) // もう既に1回関数を呼び出したので、ここは実行されない！

	fmt.Printf("Count: %d\n", count)
}
