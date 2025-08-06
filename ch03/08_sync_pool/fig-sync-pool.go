package main

import (
	"fmt"
	"sync"
)

// main関数はプログラムのエントリーポイントで、sync.Poolを使用したメモリ管理と並行処理の実行を行います。
// sync.Poolを利用して効率的にメモリバッファを再利用する仕組みを構築します。
// 各ゴルーチンはsync.Poolからメモリバッファを取得し、処理後にバッファを返却します。
// メモリバッファの生成数を追跡し、最終的に標準出力に出力します。
// プールに初期データを設定し、並行処理による高負荷に対応できるようにします。
func main() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem // <1>
		},
	}

	// Seed the pool with 4KB
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()

			mem := calcPool.Get().(*[]byte) // <2>
			defer calcPool.Put(mem)

			// Assume something interesting, but quick is being done with
			// this memory.
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)
}
