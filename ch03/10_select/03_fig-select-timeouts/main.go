package main

import (
	"fmt"
	"time"
)

// main関数はプログラムのエントリーポイントで、select文でタイムアウト処理を利用します。
func main() {
	var c <-chan int
	select {
	case <-c: // このcase文はnilチャネルから読み込んでいるので決してブロックが解放されません
	case <-time.After(1 * time.Second): // 1秒後にタイムアウト
		fmt.Println("Timed out.")
	}
}
