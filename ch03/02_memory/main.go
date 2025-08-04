// Goルーチンのメモリ使用量を計測するプログラム
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c } // noop関数は何もしないが、チャネルから値を受け取ることでブロックする

	const numGoroutines = 1e4 // 1万のゴルーチンを起動する
	wg.Add(numGoroutines)
	before := memConsumed() // メモリ使用量を計測する前にGCを実行しておく
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed() // 全てのゴルーチンが終了した後のメモリ使用量を計測する
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}
