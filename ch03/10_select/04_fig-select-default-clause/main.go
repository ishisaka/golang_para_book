package main

import (
	"fmt"
	"time"
)

// main はプログラムのエントリーポイントとなり、select構文を使用した非同期処理の動作を確認します。
func main() {
	start := time.Now()
	var c1, c2 <-chan int
	select {
	case <-c1:
	case <-c2:
	default:
		// ブロッキングされているc1,c2の受信を待たずに以下が実行される。
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}
}
