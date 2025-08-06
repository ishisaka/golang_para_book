// fig_sync_once_do_deadlock.go
package main

import (
	"sync"
)

// 循環参照によりデッドロックしてしまう例
func main() {
	var onceA, onceB sync.Once
	var initB func()
	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) } // この関数呼び出しは<2>の関数呼び出しが値を返すまで起こらない
	onceA.Do(initA)                    // <2>
}
