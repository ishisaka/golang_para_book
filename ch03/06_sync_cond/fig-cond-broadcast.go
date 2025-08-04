// fig-cond-broadcast.go
// このコードは、sync.Condを使用して、条件に基づくイベントのブロードキャストを
// 行うサンプルプログラムです。
package main

import (
	"fmt"
	"sync"
)

func main() {
	type Button struct { // Clickedという条件を含んでいるButton型を定義します。
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) { // 条件にサブスクライブする関数を定義します。
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup // クリックイベントの登録を待つためのWaitGroupを定義します。
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() { // ボタンクリックのハンドラーを登録します
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() { // ボタンクリックのハンドラーを登録します
		fmt.Println("Displaying annoying dialogue box!")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() { // ボタンクリックのハンドラーを登録します
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast() // クリックイベントをブロードキャストします。

	clickRegistered.Wait()
}
