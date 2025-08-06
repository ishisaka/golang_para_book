package main

import (
	"fmt"
	"sync"
)

// main はプログラムのエントリーポイントであり、sync.Poolの基本的な使用例を示します。
func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	myPool.Get()             // Newを呼び出す
	instance := myPool.Get() // Newを呼び出す
	myPool.Put(instance)     // 作成したインスタンスをPoolに戻す
	myPool.Get()             // プールされたインスタンスを呼び出す
}
