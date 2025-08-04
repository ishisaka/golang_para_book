package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // WaitGroupのカウンタを1に設定
	go func() {
		defer wg.Done() // 処理が終わったらカウンタを1減らす
		fmt.Println("hello")
	}()
	// 他の処理を続ける
	// time.Sleep(2000 * time.Millisecond) // SleepでGo routinesの終了を待つのは新たな競合を作ってしまっている
	// 代わりにWaitGroupを使い、fork-joinモデルを実装してGo routinesの終了を待つ。
	wg.Wait() // WaitGroupのカウンタが0になるまで待つ

	// クロージャーを使った例
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) { // クロージャーを使って変数をキャプチャ
			defer wg.Done()
			fmt.Println(salutation) // キャプチャした変数を使って出力
		}(salutation) // 引数として渡すことで、ループ変数の値を正しくキャプチャ
	}
	wg.Wait()
}
