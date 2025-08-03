/*
これはデータ競合は解決したが、データ競合状態は解決されていない。
*/
package main

import (
	"sync"
)

func main() {
	var memoryAccess sync.Mutex
	var data int
	go func() {
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if data == 0 {
		println("data is still 0")
	} else {
		println("data has been incremented to", data)
	}
	memoryAccess.Unlock()
}
