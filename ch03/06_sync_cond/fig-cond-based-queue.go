// fig-cond-based-queue.go
// このコードは、sync.Condを使用して、条件に基づくキューの実装を行うサンプルプログラムです。
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})    // 標準のsync.MutexをLockerとして使って条件を作成します
	queue := make([]interface{}, 0, 10) // 長さ0のスライスを作成します。

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()        // 再度条件のクリティカルセクションに入って、条件に合った形でデータを修正します。
		queue = queue[1:] // キューの先頭要素を削除します。
		fmt.Println("Removed from queue")
		c.L.Unlock() // クリティカルセクションを抜けます。
		c.Signal()   // 待機中のゴルーチンにシグナルを送ります。
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()            // 条件であるLockerのLockメソッドを呼び出してクリティカルセクションに入ります。
		for len(queue) == 2 { // キューの長さを確認します。
			c.Wait() // 条件のシグナルが送出されるまでメインゴルーチンを一時停止します。
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second) // 1秒後に要素をキューから取り出す新しいゴルーチンを生成します。
		c.L.Unlock()                        // 要素を無事キューに追加できたので条件のクリティカルセクションを抜けます。
	}
}
