package main

import (
	"fmt"
)

// チャネルのオーナーシップの例
// チャネルの初期化、書き込み、チャネルのクローズをカプセル化することでnilのチャネルへのアクセスを防ぐ
func main() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5) // バッファ付きチャネルを初期化
		go func() {                       // resultStreamへ書き込むゴールーチンを起動
			defer close(resultStream) // 使い終わったらチャネルを閉じる
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream // チャネルを返す
	}

	resultStream := chanOwner()
	for result := range resultStream { // チャネルから値を取得する
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}
