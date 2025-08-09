package main

import (
	"fmt"
)

// レキシカル拘束の例
// チャネルをレキシカルスコープの中に閉じ込めることで不意なチャネルへのアクセスを防ぐ
func main() {
	chanOwner := func() <-chan int {
		results := make(chan int, 5) // チャネルをchanOwner関数のレキシカルスコープ内で初期化します。これによってresultsチャネルへの書き込みができるスコープを制限しています。
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) { // intのチャネルの読み込み専用のコピーを受け取ります
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner() // チャネルへの読み込み権限を受け取って、消費者に渡しています。
	consumer(results)
}
